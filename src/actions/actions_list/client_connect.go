package actionslist

import (
	log "github.com/sirupsen/logrus"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func ClientConnect(action_params []string, context *models.Context) {
	log.Debugf("Client Connect: %v", action_params)
}