package commandslist

import (
	goto_shared "github.com/AntoineMeresse/flibot-urt/src/commands/shared/goto"
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func SetGoto(cmd *models.CommandsArgs) {
	if len(cmd.Params) > 0 {
		jumpname := goto_shared.GetJumpNameForSavePos(cmd.Server, cmd.Params[0])
		cmd.RconCommand("saveJumpPos %s %s", cmd.PlayerId, jumpname)
	}
}