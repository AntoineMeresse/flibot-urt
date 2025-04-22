package actionslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func ClientSpawn(actionParams []string, _ *context.Context) {
	// When the player join in game
	log.Debugf("ClientSpawn: %v", actionParams)
}
