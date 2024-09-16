package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Quit(cmd *models.CommandsArgs) {
	cmd.RconCommand("kick %s \"%s\"", cmd.PlayerId, "This player isn't good enough for this map !")
}