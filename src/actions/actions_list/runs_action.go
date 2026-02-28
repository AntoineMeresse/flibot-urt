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
	playerNumber := actionParams[0]
	wayName := actionParams[3]

	var bestCheckpoints []int
	var bestPlayerName string
	if cps, name, err := c.DB.GetBestCheckpoints(c.GetCurrentMap(), wayName); err == nil {
		bestCheckpoints = cps
		bestPlayerName = name
	}
	c.Runs.RunStart(playerNumber, wayName, bestCheckpoints, bestPlayerName)
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
	playerNumber := actionParams[0]
	c.Runs.AddCheckpoint(playerNumber, actionParams[6])

	if c.Runs.IsCpEnabled(playerNumber) {
		if player, err := c.Players.GetPlayer(playerNumber); err == nil {
			if msg := c.Runs.GetCpMsg(playerNumber, player.Name); msg != "" {
				c.RconText(false, playerNumber, "%s", msg)
			}
		}
	}
}

func RunLog(actionParams []string, c *appcontext.AppContext) {
	runJson := strings.Join(actionParams, " ")
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
			if isImprovement {
				c.GivePenCoin(*player)
			}
			go func() {
				sendToDiscordWebhook(c, runInfo, discordMsg)
				moveDemoIfImprovement(c, runInfo, isImprovement)
			}()
		}
	}
}

func processRunData(c *appcontext.AppContext, r api.SendDemoResponse, playerNumber string) (string, bool) {
	log.Debugf("SendDemoResponse: %+v", r)
	gameMsg := ""
	discordMsg := ""
	isImprovement := false

	cleanDiff := func(v string) string {
		stripped := strings.TrimPrefix(v, "-")
		if utils.IsZero(stripped) {
			return stripped
		}
		return v
	}
	diffColor := func(v string) string {
		if strings.Contains(v, "-") && !utils.IsZero(strings.TrimPrefix(v, "-")) {
			return "^1"
		}
		return "^2"
	}

	if r.Improvement != "" {
		discordMsg += fmt.Sprintf("PB difference: %s", r.Improvement)
		if utils.IsImprovement(r.Improvement) {
			isImprovement = true
		}
		gameMsg += fmt.Sprintf("^5PB ^7difference: %s%s^7", diffColor(r.Improvement), cleanDiff(r.Improvement))
	}

	if r.Wrdifference != "" {
		if gameMsg != "" {
			gameMsg += " | "
			discordMsg += " | "
		}
		if utils.IsImprovement(r.Wrdifference) {
			isImprovement = true
			discordMsg += fmt.Sprintf("WR difference: %s. New WR, gg!", r.Wrdifference)
		} else {
			discordMsg += fmt.Sprintf("WR difference: %s", r.Wrdifference)
		}
		gameMsg += fmt.Sprintf("^5WR ^7difference: %s%s^7", diffColor(r.Wrdifference), cleanDiff(r.Wrdifference))
	}

	if r.Rank != nil {
		discordMsg += fmt.Sprintf(" (Rank: %s)", *r.Rank)
		gameMsg += fmt.Sprintf(" ^7(^3%s^7)", *r.Rank)
	}

	c.RconText(isImprovement, playerNumber, gameMsg)

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

func moveDemoIfImprovement(c *appcontext.AppContext, runInfo models.PlayerRunInfo, isImprovement bool) {
	if !isImprovement {
		return
	}
	src := c.UrtConfig.DemoPath + "/" + runInfo.GetDemoName()
	destDir := c.UrtConfig.DemoPath + "/improvement"
	if err := os.MkdirAll(destDir, 0755); err != nil {
		log.Errorf("Failed to create improvement directory: %v", err)
		return
	}
	dest := destDir + "/" + runInfo.GetDemoName()
	if err := os.Rename(src, dest); err != nil {
		log.Errorf("Failed to move improvement demo: %v", err)
	} else {
		log.Debugf("Moved improvement demo to: %s", dest)
	}
}
