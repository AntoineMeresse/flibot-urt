package actionslist

import (
	"encoding/json"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	log "github.com/sirupsen/logrus"
	"strings"
)

func ClientJumpRunStarted(actionParams []string, context *models.Context) {
	log.Debugf("ClientJumpRunStarted: %v", actionParams)
	if len(actionParams) < 4 {
		log.Error("ClientJumpRunStarted: Invalid parameters")
		return
	}
	context.Runs.RunStart(actionParams[0], actionParams[3])
}

func ClientJumpRunCanceled(actionParams []string, context *models.Context) {
	log.Debugf("ClientJumpRunCanceled: %v", actionParams)
	context.Runs.RunCanceled(actionParams[0])
}

func ClientJumpRunStopped(actionParams []string, context *models.Context) {
	log.Debugf("ClientJumpRunStopped: %v", actionParams)
	if len(actionParams) < 7 {
		log.Error("ClientJumpRunStopped: Invalid parameters")
		return
	}
	if player, err := context.Players.GetPlayer(actionParams[0]); err == nil {
		context.Runs.RunStopped(actionParams[0], player.Guid, actionParams[6])
	}
}

func ClientJumpRunCheckpoint(actionParams []string, context *models.Context) {
	log.Debugf("ClientJumpRunCheckpoint: %v", actionParams)
	if len(actionParams) < 7 {
		log.Error("ClientJumpRunCheckpoint: Invalid parameters")
		return
	}
	context.Runs.AddCheckpoint(actionParams[0], actionParams[6])
}

func RunLog(actionParams []string, context *models.Context) {
	runJson := strings.Join(actionParams, "")
	runJson = strings.Replace(runJson, "'", "\"", -1)
	log.Debugf("RunLog: %v", runJson)

	var runInfo PlayerRunInfo
	if err := json.Unmarshal([]byte(runJson), &runInfo); err != nil {
		log.Errorf("RunLog: Error unmarshalling json: %v", err)
	} else {
		if player, err := context.Players.GetPlayer(runInfo.Playernumber); err == nil {
			cps := context.Runs.RunGetCheckpoint(player.Id, player.Guid, runInfo.Time, runInfo.Way)
			//TODO: Send to ujm
			context.RconText(false, runInfo.Playernumber, "%s: %s (%v)", runInfo.Playername, runInfo.Time, cps)
		}
	}
}

type PlayerRunInfo struct {
	Server       string `json:"server"`
	ServerName   string `json:"server_name"`
	Fps          string `json:"fps"`
	Mapname      string `json:"mapname"`
	Playername   string `json:"playername"`
	Guid         string `json:"guid"`
	Way          string `json:"way"`
	Time         string `json:"time"`
	Demopath     string `json:"demopath"`
	Playernumber string `json:"playernumber"`
	GUtj         string `json:"g_utj"`
}
