package commandslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Stamina(cmd *models.CommandsArgs) {
	cmd.RconCommand(fmt.Sprintf("customstamina %s", cmd.PlayerId))
}