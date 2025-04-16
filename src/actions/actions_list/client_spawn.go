package actionslist

import (
	log "github.com/sirupsen/logrus"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func ClientSpawn(actionParams []string, context *models.Context) {
	// When the player join in game
	log.Debugf("ClientSpawn: %v", actionParams)
}
