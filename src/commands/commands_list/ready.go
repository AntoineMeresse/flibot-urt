package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Ready(cmd *models.CommandsArgs) {
	cmd.RconCommand("ready %s", cmd.PlayerId)
}