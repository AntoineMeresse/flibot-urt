package actionslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
	log "github.com/sirupsen/logrus"
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
	for k, v := range actionParams {
		log.Debugf("%d -> %s", k, v)
	}

	log.Debugf("ClientJumpRunCanceled: %v", actionParams)
}

func ClientJumpRunStopped(actionParams []string, context *models.Context) {
	log.Debugf("ClientJumpRunStopped: %v", actionParams)
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
	log.Debugf("RunLog: %v", actionParams)
}
