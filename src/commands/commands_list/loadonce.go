package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Loadonce(cmd *models.CommandsArgs) {
	cmd.RconCommand("simpleload %s", cmd.PlayerId)
}