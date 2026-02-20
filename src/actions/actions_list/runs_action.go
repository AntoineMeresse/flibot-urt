package actionslist

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/AntoineMeresse/flibot-urt/src/api"
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func ClientJumpRunStarted(actionParams []string, c *appcontext.AppContext) {
	log.Debugf("ClientJumpRunStarted: %v", actionParams)
	if len(actionParams) < 4 {
		log.Error("ClientJumpRunStarted: Invalid parameters")
		return
	}
	c.Runs.RunStart(actionParams[0], actionParams[3])
}

func ClientJumpRunCanceled(actionParams []string, c *appcontext.AppContext) {
	log.Debugf("ClientJumpRunCanceled: %v", actionParams)
	if len(actionParams) > 0 {
		c.Runs.RunCanceled(actionParams[0])
	}
}

func ClientJumpRunStopped(actionParams []string, c *appcontext.AppContext) {
	log.Debugf("ClientJumpRunStopped: %v", actionParams)
	if len(actionParams) < 7 {
		log.Error("ClientJumpRunStopped: Invalid parameters")
		return
	}
	if player, err := c.Players.GetPlayer(actionParams[0]); err == nil {
		c.Runs.RunStopped(actionParams[0], player.Guid, actionParams[6])
	}
}

func ClientJumpRunCheckpoint(actionParams []string, c *appcontext.AppContext) {
	log.Debugf("ClientJumpRunCheckpoint: %v", actionParams)
	if len(actionParams) < 7 {
		log.Error("ClientJumpRunCheckpoint: Invalid parameters")
		return
	}
	c.Runs.AddCheckpoint(actionParams[0], actionParams[6])
}

func RunLog(actionParams []string, c *appcontext.AppContext) {
	runJson := strings.Join(actionParams, "")
	runJson = strings.Replace(runJson, "'", "\"", -1)
	log.Debugf("RunLog: %v", runJson)

	var runInfo models.PlayerRunInfo
	if err := json.Unmarshal([]byte(runJson), &runInfo); err != nil {
		log.Errorf("RunLog: Error unmarshalling json: %v", err)
	} else {
		if player, err := c.Players.GetPlayer(runInfo.Playernumber); err == nil {
			cps := c.Runs.RunGetCheckpoint(player.Number, player.Guid, runInfo.Time, runInfo.Way)
			runInfo.PlayerIp = player.Ip
			if runInfo.PlayerIp == "" {
				if dbPlayer, found := c.DB.GetPlayerByGuid(player.Guid); found {
					runInfo.PlayerIp = dbPlayer.Ip
				}
			}

			if err := c.DB.HandleRun(runInfo, cps); err != nil {
				log.Errorf("RunLog: Error handling run: %v", err)
			}

			var demoResponse api.SendDemoResponse
			if runInfo.Utj == "0" {
				demoResponse, err = c.Api.PostRunDemo(runInfo, cps, c.UrtConfig.DemoPath)
				if err != nil {
					log.Errorf("RunLog: Error posting run: %v", err)
				}
			}
			processRunData(c, demoResponse, player.Number)
			go func() {
				msg := fmt.Sprintf("[Flibot] %s finished %s of %s in %s.", runInfo.Playername, runInfo.Way,
					runInfo.Mapname, runInfo.Time)
				if err := c.Api.SendFileToWebhook(runInfo.GetDemoName(), msg); err != nil {
					log.Errorf("Webhook send failed: %v", err)
				}
			}()
		}
	}
}

func processRunData(c *appcontext.AppContext, r api.SendDemoResponse, playerNumber string) {
	logrus.Debugf("SendDemoResponse: %+v", r)
	// discordMsg := "discord: "
	ingameMsg := ""
	global := false

	if r.Improvement != "" {
		if utils.IsImprovement(r.Improvement) {
			ingameMsg += fmt.Sprintf("^5PB ^7difference: ^2%s^7", r.Improvement)
			global = true
		} else {
			ingameMsg += fmt.Sprintf("^5PB ^7difference: ^1%s^7", r.Improvement)
		}
	}

	if r.Wrdifference != "" {
		if ingameMsg != "" {
			ingameMsg += " | "
		}

		if utils.IsImprovement(r.Wrdifference) {
			global = true
			ingameMsg += fmt.Sprintf("^5WR ^7difference: ^2%s^7", r.Wrdifference)
		} else {
			ingameMsg += fmt.Sprintf("^5WR ^7difference: ^1%s^7", r.Wrdifference)
		}
	}

	if r.Rank != nil {
		ingameMsg += fmt.Sprintf("  ^7(^3%s^7)", *r.Rank)
	}

	c.RconText(global, playerNumber, ingameMsg)
}
