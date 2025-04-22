package commandslist

import (
	gotoshared "github.com/AntoineMeresse/flibot-urt/src/commands/shared/goto"
	"github.com/AntoineMeresse/flibot-urt/src/context"
)

func SetGoto(cmd *context.CommandsArgs) {
	if len(cmd.Params) > 0 {
		jumpname := gotoshared.GetJumpNameForSavePos(cmd.Context, cmd.Params[0])
		cmd.RconCommand("saveJumpPos %s %s", cmd.PlayerId, jumpname)
	}
}
