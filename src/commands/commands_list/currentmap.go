package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func CurrentMap(cmd *models.CommandsArgs) {
	cmd.RconText(cmd.Context.GetCurrentMap())
}