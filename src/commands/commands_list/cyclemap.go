package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func CycleMap(cmd *appcontext.CommandsArgs) {
	_, force := utils.ExtractForceFlag(cmd.Params)

	if n := cmd.Context.Runs.AnyRunning(); n > 0 && !force {
		cmd.RconText("^3%d^7 player(s) are currently running. Add ^3-f^7 to cycle anyway.", n)
		return
	}

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
