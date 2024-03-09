package commandslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Invisible(server *models.Server, playerNumber string, params []string, isGlobal bool) {
	server.Rcon.RconCommand(fmt.Sprintf("invisible %s", playerNumber))
}