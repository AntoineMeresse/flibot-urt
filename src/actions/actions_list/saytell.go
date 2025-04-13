package actionslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
	log "github.com/sirupsen/logrus"
)

func SayTell(actionParams []string, context *models.Context) {
	log.Debugf("SayTell: %v", actionParams)
}
