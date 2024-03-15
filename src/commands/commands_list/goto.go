package commandslist

import (
	"fmt"

	goto_shared "github.com/AntoineMeresse/flibot-urt/src/commands/shared/goto"
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Goto(cmd *models.CommandsArgs) {
	if len(cmd.Params) == 0 {
		locationDisplayList := goto_shared.GetDisplayLocation(cmd.Server)
		cmd.RconList(locationDisplayList)
	} else {
		jumpName := cmd.Params[0]
		if exists, _  := goto_shared.DoesPositionExist(cmd.Server, jumpName); exists {
			cmd.RconCommand(fmt.Sprintf("forceteam %s free", cmd.PlayerId))
			cmd.RconCommand(fmt.Sprintf("loadJumpPos %s %s", cmd.PlayerId, jumpName))
		} else {
			cmd.RconText(fmt.Sprintf("Location (^5%s^3) ^1doesn't^3 exist.", jumpName))
		}
	}
}