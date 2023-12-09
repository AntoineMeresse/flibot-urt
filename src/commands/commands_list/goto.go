package commandslist

import (
	"fmt"

	goto_shared "github.com/AntoineMeresse/flibot-urt/src/commands/shared/goto"
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Goto(server models.Server, playerNumber string, params []string, isGlobal bool) {
	if len(params) == 0 {
		locationDisplayList := goto_shared.GetDisplayLocation(server)
		server.RconList(locationDisplayList, isGlobal, playerNumber)
	} else {
		jumpName := params[0]
		if goto_shared.DoesPositionExist(server, jumpName) {
			server.Rcon.RconCommand(fmt.Sprintf("forceteam %s free", playerNumber))
			server.Rcon.RconCommand(fmt.Sprintf("loadJumpPos %s %s", playerNumber, jumpName))
		} else {
			server.RconText(fmt.Sprintf("^3Location (^5%s^3) ^1doesn't^3 exist.", jumpName), isGlobal, playerNumber)
		}
	}
}