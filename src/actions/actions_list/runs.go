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
	log.Debugf("ClientJumpRunCanceled: %v", actionParams)
}

func ClientJumpRunStopped(actionParams []string, context *models.Context) {
	log.Debugf("ClientJumpRunStopped: %v", actionParams)
}

func ClientJumpRunCheckpoint(actionParams []string, context *models.Context) {
	log.Debugf("ClientJumpRunCheckpoint: %v", actionParams)

	for k, v := range actionParams[2:] {
		log.Debugf("%d -> %s", k, v)
	}
}

func RunLog(actionParams []string, context *models.Context) {
	log.Debugf("RunLog: %v", actionParams)
}
