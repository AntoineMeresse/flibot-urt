package actionslist

import (
	"fmt"
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func Radio(actionParams []string, c *appcontext.AppContext) {
	log.Debugf("Radio: %v", actionParams)
	full := strings.Join(actionParams, " ")
	parts := strings.SplitN(full, "-", 4)
	if len(parts) < 4 {
		log.Warnf("Radio: unexpected format: %s", full)
		return
	}
	playerNum := strings.TrimSpace(parts[0])
	r1 := parts[1]
	r2 := parts[2]
	rest := strings.TrimSpace(strings.ReplaceAll(parts[3], `""`, ""))

	player, err := c.Players.GetPlayer(playerNum)
	if err != nil {
		log.Warnf("Radio: player not found: %s", playerNum)
		return
	}

	team := "[game] "
	if player.IsSpec() {
		team = "[spec] "
	}
	c.SendBridgeMessage(fmt.Sprintf("%s: Radio [%s%s] %s", player.Name, r1, r2, rest), team)
}
