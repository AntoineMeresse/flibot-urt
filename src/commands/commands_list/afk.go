package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Afk(cmd *appcontext.CommandsArgs) {
	args, force := utils.ExtractForceFlag(cmd.Params)
	if len(args) > 0 {
		player, err := cmd.Context.Players.GetPlayer(args[0])
		if err == nil {
			if cmd.Context.Runs.IsRunning(player.Number) && !force {
				cmd.RconText("^5%s^3 is currently running. Add ^3-f^7 to move to spec anyway.", player.Name)
				return
			}
			cmd.RconCommand("forceteam %s spec", player.Number)
		} else {
			cmd.RconText(err.Error())
		}
	}
}
