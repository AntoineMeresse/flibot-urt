package actionslist

import (
	log "github.com/sirupsen/logrus"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func ClientBegin(action_params []string, context *models.Context) {
	log.Debugf("Client Begin: %v", action_params)
}