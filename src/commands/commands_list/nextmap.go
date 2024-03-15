package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func NextMap(cmd *models.CommandsArgs) {
	cmd.RconText(cmd.Server.GetNextMap())
}