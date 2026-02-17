package commandslist

import (
	gotoshared "github.com/AntoineMeresse/flibot-urt/src/commands/shared/goto"
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils/msg"
)

func Goto(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		locationDisplayList := gotoshared.GetDisplayLocation(cmd.Context)
		cmd.RconList(locationDisplayList)
	} else {
		jumpName := cmd.Params[0]
		pos, err := cmd.Context.DB.PositionGet(cmd.Context.GetCurrentMap(), jumpName)
		if err != nil {
			cmd.RconText(msg.GOTO_NO_LOCATION, jumpName)
			return
		}
		cmd.RconCommand("forceteam %s free", cmd.PlayerId)
		cmd.RconCommand("goToPosition %s %g %g %g %g", cmd.PlayerId, pos.X, pos.Y, pos.Z, pos.Angle)
	}
}
