package actionslist

import (
	"fmt"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func ClientDisconnect(actionParams []string, c *appcontext.AppContext) {
	log.Debugf("ClientDisconnect: %v", actionParams)
	log.Debugf("[Before] ClientDisconnect: %v", c.Players.PlayerMap)
	if len(actionParams) > 0 {
		playerNumber := actionParams[0]
		if player, err := c.Players.GetPlayer(playerNumber); err == nil {
			c.SendBridgeMessage(fmt.Sprintf("%s has left the game. Bye !", player.Name), "")
		}
		c.Runs.RunCanceled(playerNumber)
		c.ClearTrad(playerNumber)
		c.Players.RemovePlayer(playerNumber)
	}
	log.Debugf("[After] ClientDisconnect: %v", c.Players.PlayerMap)
}
