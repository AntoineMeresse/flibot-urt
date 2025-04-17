package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Afk(cmd *models.CommandsArgs) {
	if len(cmd.Params) > 0 {
		player, err := cmd.Context.Players.GetPlayer(cmd.Params[0])
		if err == nil {
			// TODO: Check if player isn't running.
			cmd.RconCommand("forceteam %s spec", player.Number)
		} else {
			cmd.RconText(err.Error())
		}
	}
}
