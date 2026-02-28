package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func MovePlayer(cmd *appcontext.CommandsArgs) {
	args, force := utils.ExtractForceFlag(cmd.Params)
	if len(args) < 2 {
		cmd.RconUsage()
		return
	}

	p1, err := cmd.Context.Players.GetPlayer(args[0])
	if err != nil {
		cmd.RconText("%s", err.Error())
		return
	}

	p2, err := cmd.Context.Players.GetPlayer(args[1])
	if err != nil {
		cmd.RconText("%s", err.Error())
		return
	}

	if p1.Number == p2.Number {
		cmd.RconText("^7Both players are the same.")
		return
	}

	if cmd.Context.Runs.IsRunning(p1.Number) && !force {
		cmd.RconText("^5%s^3 is currently running. Add ^3-f^7 to move anyway.", p1.Name)
		return
	}

	cmd.RconCommand("forceteam %s free", p1.Number)
	cmd.RconCommand("tp %s %s", p1.Number, p2.Number)
}
