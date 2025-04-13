package actionslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
	log "github.com/sirupsen/logrus"
)

func Tell(actionParams []string, context *models.Context) {
	log.Debugf("Tell: %v", actionParams)
}
