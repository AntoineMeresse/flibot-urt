package commandslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Play(server *models.Server, playerNumber string, params []string, isGlobal bool) {
	server.Rcon.RconCommand(fmt.Sprintf("forceteam %s free", playerNumber))
}