package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Invisible(cmd *models.CommandsArgs) {
	cmd.RconCommand("invisible %s", cmd.PlayerId)
}