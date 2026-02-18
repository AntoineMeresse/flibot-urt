package actionslist

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/AntoineMeresse/flibot-urt/src/api"
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func ClientJumpRunStarted(actionParams []string, c *appcontext.AppContext) {
	slog.Debug("ClientJumpRunStarted", "params", actionParams)
	if len(actionParams) < 4 {
		slog.Error("ClientJumpRunStarted: invalid parameters")
		return
	}
	c.Runs.RunStart(actionParams[0], actionParams[3])
}

func ClientJumpRunCanceled(actionParams []string, c *appcontext.AppContext) {
	slog.Debug("ClientJumpRunCanceled", "params", actionParams)
	if len(actionParams) > 0 {
		c.Runs.RunCanceled(actionParams[0])
	}
}

func ClientJumpRunStopped(actionParams []string, c *appcontext.AppContext) {
	slog.Debug("ClientJumpRunStopped", "params", actionParams)
	if len(actionParams) < 7 {
		slog.Error("ClientJumpRunStopped: invalid parameters")
		return
	}
	if player, err := c.Players.GetPlayer(actionParams[0]); err == nil {
		c.Runs.RunStopped(actionParams[0], player.Guid, actionParams[6])
	}
}

func ClientJumpRunCheckpoint(actionParams []string, c *appcontext.AppContext) {
	slog.Debug("ClientJumpRunCheckpoint", "params", actionParams)
	if len(actionParams) < 7 {
		slog.Error("ClientJumpRunCheckpoint: invalid parameters")
		return
	}
	c.Runs.AddCheckpoint(actionParams[0], actionParams[6])
}

func RunLog(actionParams []string, c *appcontext.AppContext) {
	runJson := strings.Join(actionParams, "")
	runJson = strings.Replace(runJson, "'", "\"", -1)
	slog.Debug("RunLog", "json", runJson)

	var runInfo models.PlayerRunInfo
	if err := json.Unmarshal([]byte(runJson), &runInfo); err != nil {
		slog.Error("RunLog: failed to unmarshal json", "err", err)
	} else {
		if player, err := c.Players.GetPlayer(runInfo.Playernumber); err == nil {
			cps := c.Runs.RunGetCheckpoint(player.Number, player.Guid, runInfo.Time, runInfo.Way)
			runInfo.PlayerIp = player.Ip

			if err := c.DB.HandleRun(runInfo, cps); err != nil {
				slog.Error("RunLog: error handling run", "err", err)
			}

			var demoResponse api.SendDemoResponse
			if runInfo.Utj == "0" {
				var demoErr error
				demoResponse, demoErr = c.Api.PostRunDemo(runInfo, c.UrtConfig.DemoPath)
				if demoErr != nil {
					slog.Error("RunLog: error posting run", "err", demoErr)
				}
			}
			processRunData(c, demoResponse, player.Number)
			go func() {
				msg := fmt.Sprintf("[Flibot] %s finished %s of %s in %s.", runInfo.Playername, runInfo.Way,
					runInfo.Mapname, runInfo.Time)
				if err := c.Api.SendFileToWebhook(runInfo.GetDemoName(), msg); err != nil {
					slog.Error("Webhook send failed", "err", err)
				}
			}()
		}
	}
}

func processRunData(c *appcontext.AppContext, r api.SendDemoResponse, playerNumber string) {
	slog.Debug("SendDemoResponse", "response", r)
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
