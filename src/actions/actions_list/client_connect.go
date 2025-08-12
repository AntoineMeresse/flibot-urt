package actionslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func ClientConnect(actionParams []string, c *appcontext.AppContext) {
	log.Debugf("Client Connect: %v", actionParams)
	c.Players.AddPlayer(actionParams[0], &models.Player{})
}
