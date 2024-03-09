package commandslist

import (
	"fmt"

	goto_shared "github.com/AntoineMeresse/flibot-urt/src/commands/shared/goto"
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func RemoveGoto(server *models.Server, playerNumber string, params []string, isGlobal bool) {
	if len(params) > 0 {
		jumpName := params[0]
		deleted := goto_shared.RemovePosition(server, jumpName)
		if deleted {
			server.RconText(fmt.Sprintf("Location (^5%s^3) has been deleted.", jumpName), isGlobal, playerNumber)
		} else {
			server.RconText(fmt.Sprintf("Location (^5%s^3) ^1doesn't^3 exist.", jumpName), isGlobal, playerNumber)
		}
	}
}