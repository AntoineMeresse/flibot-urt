package commandslist

import (
	gotoshared "github.com/AntoineMeresse/flibot-urt/src/commands/shared/goto"
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils/msg"
)

func RemoveGoto(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) > 0 {
		jumpName := cmd.Params[0]
		deleted := gotoshared.RemovePosition(cmd.Context, jumpName)
		if deleted {
			cmd.RconText(msg.GOTO_REMOVE, jumpName)
		} else {
			cmd.RconText(msg.GOTO_DONT_EXIST, jumpName)
		}
	}
}
