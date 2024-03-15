package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Stamina(cmd *models.CommandsArgs) {
	cmd.RconCommand("customstamina %s", cmd.PlayerId)
}