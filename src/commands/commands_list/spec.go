package commandslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Spec(cmd *models.CommandsArgs) {
	cmd.RconCommand(fmt.Sprintf("forceteam %s spec", cmd.PlayerId))
}