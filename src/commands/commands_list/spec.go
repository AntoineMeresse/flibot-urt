package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Spec(cmd *models.CommandsArgs) {
	cmd.RconCommand("forceteam %s spec", cmd.PlayerId)
}