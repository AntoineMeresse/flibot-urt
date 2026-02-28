package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Tp(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconUsage()
		return
	}

	args, force := utils.ExtractForceFlag(cmd.Params)
	if len(args) == 0 {
		cmd.RconUsage()
		return
	}

	target, err := cmd.Context.Players.GetPlayer(args[0])
	if err != nil {
		cmd.RconText("%s", err.Error())
		return
	}

	if target.Number == cmd.PlayerId {
		cmd.RconText("^7You cannot teleport to yourself.")
		return
	}

	if cmd.Context.Runs.IsRunning(cmd.PlayerId) && !force {
		cmd.RconText("^3You are currently running. Add ^3-f^7 to teleport anyway.")
		return
	}

	cmd.RconCommand("forceteam %s free", cmd.PlayerId)
	cmd.RconCommand("tp %s %s", cmd.PlayerId, target.Number)
}
