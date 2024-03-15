package commandslist

import (
	goto_shared "github.com/AntoineMeresse/flibot-urt/src/commands/shared/goto"
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func RemoveGoto(cmd *models.CommandsArgs) {
	if len(cmd.Params) > 0 {
		jumpName := cmd.Params[0]
		deleted := goto_shared.RemovePosition(cmd.Server, jumpName)
		if deleted {
			cmd.RconText("Location (^5%s^3) has been deleted.", jumpName)
		} else {
			cmd.RconText("Location (^5%s^3) ^1doesn't^3 exist.", jumpName)
		}
	}
}