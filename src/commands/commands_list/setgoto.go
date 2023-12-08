package commandslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func SetGoto(server models.Server, playerNumber string, params []string, isGlobal bool) {
	if len(params) > 0 {
		jumpname := params[0]
		server.Rcon.RconCommand(fmt.Sprintf("saveJumpPos %s %s", playerNumber, jumpname))
	}
}