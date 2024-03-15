package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Play(cmd *models.CommandsArgs) {
	cmd.RconCommand("forceteam %s free", cmd.PlayerId)
}