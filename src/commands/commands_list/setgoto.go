package commandslist

import (
	"fmt"

	goto_shared "github.com/AntoineMeresse/flibot-urt/src/commands/shared/goto"
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func SetGoto(server models.Server, playerNumber string, params []string, isGlobal bool) {
	if len(params) > 0 {
		jumpname := goto_shared.GetJumpNameForSavePos(server, params[0])
		server.Rcon.RconCommand(fmt.Sprintf("saveJumpPos %s %s", playerNumber, jumpname))
	}
}