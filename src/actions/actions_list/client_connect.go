package actionslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func ClientConnect(actionParams []string, _ *appcontext.AppContext) {
	log.Debugf("Client Connect: %v", actionParams)
	//c.Players.AddPlayer(actionParams[0], &models.Player{})
}
