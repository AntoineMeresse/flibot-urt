package commandslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Ready(cmd *models.CommandsArgs) {
	cmd.RconCommand(fmt.Sprintf("ready %s", cmd.PlayerId))
}