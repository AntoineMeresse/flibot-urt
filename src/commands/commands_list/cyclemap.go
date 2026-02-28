package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Cyclemap(cmd *appcontext.CommandsArgs) {
	player, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	if err != nil {
		cmd.RconText("^1Could not identify the player who used the command.")
		return
	}

	if player.Role <= 60 {
		v := models.Vote{Params: []string{"cycle"}, PlayerId: cmd.PlayerId}
		cmd.Context.NewVote(v)
	} else {
		cmd.RconCommand("cyclemap")
	}
}
