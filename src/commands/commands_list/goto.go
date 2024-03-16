package commandslist

import (
	goto_shared "github.com/AntoineMeresse/flibot-urt/src/commands/shared/goto"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils/msg"
)

func Goto(cmd *models.CommandsArgs) {
	if len(cmd.Params) == 0 {
		locationDisplayList := goto_shared.GetDisplayLocation(cmd.Server)
		cmd.RconList(locationDisplayList)
	} else {
		jumpName := cmd.Params[0]
		if exists, _  := goto_shared.DoesPositionExist(cmd.Server, jumpName); exists {
			cmd.RconCommand("forceteam %s free", cmd.PlayerId)
			cmd.RconCommand("loadJumpPos %s %s", cmd.PlayerId, jumpName)
		} else {
			cmd.RconText(msg.GOTO_NO_LOCATION, jumpName)
		}
	}
}