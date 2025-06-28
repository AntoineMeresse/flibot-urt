package actionslist

import (
	"encoding/json"
	"github.com/AntoineMeresse/flibot-urt/src/api"
	"github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	log "github.com/sirupsen/logrus"
	"strings"
)

func ClientJumpRunStarted(actionParams []string, c *context.Context) {
	log.Debugf("ClientJumpRunStarted: %v", actionParams)
	if len(actionParams) < 4 {
		log.Error("ClientJumpRunStarted: Invalid parameters")
		return
	}
	c.Runs.RunStart(actionParams[0], actionParams[3])
}

func ClientJumpRunCanceled(actionParams []string, c *context.Context) {
	log.Debugf("ClientJumpRunCanceled: %v", actionParams)
	c.Runs.RunCanceled(actionParams[0])
}

func ClientJumpRunStopped(actionParams []string, c *context.Context) {
	log.Debugf("ClientJumpRunStopped: %v", actionParams)
	if len(actionParams) < 7 {
		log.Error("ClientJumpRunStopped: Invalid parameters")
		return
	}
	if player, err := c.Players.GetPlayer(actionParams[0]); err == nil {
		c.Runs.RunStopped(actionParams[0], player.Guid, actionParams[6])
	}
}

func ClientJumpRunCheckpoint(actionParams []string, c *context.Context) {
	log.Debugf("ClientJumpRunCheckpoint: %v", actionParams)
	if len(actionParams) < 7 {
		log.Error("ClientJumpRunCheckpoint: Invalid parameters")
		return
	}
	c.Runs.AddCheckpoint(actionParams[0], actionParams[6])
}

func RunLog(actionParams []string, c *context.Context) {
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

			if err := c.DB.HandleRun(runInfo, cps); err != nil {
				log.Errorf("RunLog: Error handling run: %v", err)
			}

			var demoResponse api.SendDemoResponse
			if runInfo.Utj == "0" {
				demoResponse, err = c.Api.PostRunDemo(runInfo, c.UrtConfig.DemoPath)
				if err != nil {
					log.Errorf("RunLog: Error posting run: %v", err)
				}
			}

			processRunData(c, demoResponse)

			c.RconText(false, runInfo.Playernumber, "%s: %s (%v)", runInfo.Playername, runInfo.Time, cps)
			c.RconText(false, runInfo.Playernumber, "Demo response: %v", demoResponse)
		}
	}
}

func processRunData(c *context.Context, r api.SendDemoResponse) {
	if r.Process {
		discordMsg := "discord: "
		ingameMsg := "ingame: "
		global := true

		if r.Improvement != "" {

		}

		c.RconText(global, "", discordMsg+" | "+ingameMsg)
	}
}
