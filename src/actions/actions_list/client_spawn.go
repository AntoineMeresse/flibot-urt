package actionslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func ClientSpawn(actionParams []string, _ *appcontext.AppContext) {
	// When the player join in game
	log.Debugf("ClientSpawn: %v", actionParams)
}
