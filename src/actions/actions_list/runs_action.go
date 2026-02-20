package actionslist

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/AntoineMeresse/flibot-urt/src/api"
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
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
	runJson = strings.ReplaceAll(runJson, "'", "\"")
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
			discordMsg, isImprovement := processRunData(c, demoResponse, player.Number)
			go func() {
				sendToDiscordWebhook(c, runInfo, discordMsg)
				deleteDemoIfNotImprovement(c, runInfo, isImprovement)
			}()
		}
	}
}

func processRunData(c *appcontext.AppContext, r api.SendDemoResponse, playerNumber string) (string, bool) {
	log.Debugf("SendDemoResponse: %+v", r)
	gameMsg := ""
	discordMsg := ""
	isImprovement := false

	if r.Improvement != "" {
		discordMsg += fmt.Sprintf("PB difference: %s", r.Improvement)
		if utils.IsImprovement(r.Improvement) {
			isImprovement = true
			gameMsg += fmt.Sprintf("^5PB ^7difference: ^2%s^7", r.Improvement)
		} else {
			gameMsg += fmt.Sprintf("^5PB ^7difference: ^1%s^7", r.Improvement)
		}
	}

	if r.Wrdifference != "" {
		if gameMsg != "" {
			gameMsg += " | "
			discordMsg += " | "
		}
		if utils.IsImprovement(r.Wrdifference) {
			isImprovement = true
			gameMsg += fmt.Sprintf("^5WR ^7difference: ^2%s^7", r.Wrdifference)
			discordMsg += fmt.Sprintf("WR difference: %s. New WR, gg!", r.Wrdifference)
		} else {
			gameMsg += fmt.Sprintf("^5WR ^7difference: ^1%s^7", r.Wrdifference)
			discordMsg += fmt.Sprintf("WR difference: %s", r.Wrdifference)
		}
	}

	if r.Rank != nil {
		discordMsg += fmt.Sprintf(" (Rank: %s)", *r.Rank)
		gameMsg += fmt.Sprintf(" ^7(^3%s^7)", *r.Rank)
	}

	if isImprovement {
		c.RconText(true, "", gameMsg)
	} else {
		c.RconText(false, playerNumber, "[PM] "+gameMsg)
	}

	return discordMsg, isImprovement
}

func sendToDiscordWebhook(c *appcontext.AppContext, runInfo models.PlayerRunInfo, discordMsg string) {
	if c.Api.DiscordWebhook == "" {
		return
	}
	msg := fmt.Sprintf("[Flibot] %s finished way %s of %s in %ss.", runInfo.Playername, runInfo.Way,
		runInfo.Mapname, utils.FormatRunTime(runInfo.Time))
	if discordMsg != "" {
		msg += fmt.Sprintf(" :stopwatch: `%s` :stopwatch:", discordMsg)
	}
	if err := c.Api.SendFileToWebhook(c.UrtConfig.DemoPath, runInfo.GetDemoName(), msg); err != nil {
		log.Errorf("Webhook send failed: %v", err)
	} else {
		log.Debugf("Demo uploaded to Discord webhook: %s", runInfo.GetDemoName())
	}
}

func deleteDemoIfNotImprovement(c *appcontext.AppContext, runInfo models.PlayerRunInfo, isImprovement bool) {
	if isImprovement {
		return
	}
	demoFile := c.UrtConfig.DemoPath + "/" + runInfo.GetDemoName()
	if err := os.Remove(demoFile); err != nil {
		log.Errorf("Failed to delete demo file: %v", err)
	} else {
		log.Debugf("Deleted non-improvement demo: %s", demoFile)
	}
}
