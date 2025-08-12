package commandslist

import (
	gotoshared "github.com/AntoineMeresse/flibot-urt/src/commands/shared/goto"
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func SetGoto(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) > 0 {
		jumpname := gotoshared.GetJumpNameForSavePos(cmd.Context, cmd.Params[0])
		cmd.RconCommand("saveJumpPos %s %s", cmd.PlayerId, jumpname)
	}
}
