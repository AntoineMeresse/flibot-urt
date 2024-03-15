package commandslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Play(cmd *models.CommandsArgs) {
	cmd.RconCommand(fmt.Sprintf("forceteam %s free", cmd.PlayerId))
}