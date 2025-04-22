package commandslist

import (
	gotoshared "github.com/AntoineMeresse/flibot-urt/src/commands/shared/goto"
	"github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils/msg"
)

func Goto(cmd *context.CommandsArgs) {
	if len(cmd.Params) == 0 {
		locationDisplayList := gotoshared.GetDisplayLocation(cmd.Context)
		cmd.RconList(locationDisplayList)
	} else {
		jumpName := cmd.Params[0]
		if exists, _ := gotoshared.DoesPositionExist(cmd.Context, jumpName); exists {
			cmd.RconCommand("forceteam %s free", cmd.PlayerId)
			cmd.RconCommand("loadJumpPos %s %s", cmd.PlayerId, jumpName)
		} else {
			cmd.RconText(msg.GOTO_NO_LOCATION, jumpName)
		}
	}
}
