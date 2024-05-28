package commandslist

import (
	goto_shared "github.com/AntoineMeresse/flibot-urt/src/commands/shared/goto"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils/msg"
)

func RemoveGoto(cmd *models.CommandsArgs) {
	if len(cmd.Params) > 0 {
		jumpName := cmd.Params[0]
		deleted := goto_shared.RemovePosition(cmd.Context, jumpName)
		if deleted {
			cmd.RconText(msg.GOTO_REMOVE, jumpName)
		} else {
			cmd.RconText(msg.GOTO_DONT_EXIST, jumpName)
		}
	}
}