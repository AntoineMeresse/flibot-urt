package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Slap(cmd *appcontext.CommandsArgs) {
	args, force := utils.ExtractForceFlag(cmd.Params)
	if len(args) == 0 {
		cmd.RconUsage()
		return
	}

	player, err := cmd.Context.Players.GetPlayer(args[0])
	if err != nil {
		cmd.RconText("%s", err.Error())
		return
	}

	if cmd.Context.Runs.IsRunning(player.Number) && !force {
		cmd.RconText("^5%s^3 is currently running. Add ^3-f^7 to slap anyway.", player.Name)
		return
	}

	cmd.RconCommand("slap %s", player.Number)
	cmd.RconText("^5%s^7 was slapped!", player.Name)
}
