package actionslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func ClientDisconnect(actionParams []string, c *appcontext.AppContext) {
	log.Debugf("ClientDisconnect: %v", actionParams)
	log.Debugf("[Before] ClientDisconnect: %v", c.Players.PlayerMap)
	if len(actionParams) > 0 {
		c.Players.RemovePlayer(actionParams[0])
	}
	log.Debugf("[After] ClientDisconnect: %v", c.Players.PlayerMap)
}
