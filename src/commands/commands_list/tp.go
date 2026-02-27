package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Tp(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconUsage()
		return
	}

	target, err := cmd.Context.Players.GetPlayer(cmd.Params[0])
	if err != nil {
		cmd.RconText("%s", err.Error())
		return
	}

	if target.Number == cmd.PlayerId {
		cmd.RconText("^7You cannot teleport to yourself.")
		return
	}

	cmd.RconCommand("forceteam %s free", cmd.PlayerId)
	cmd.RconCommand("tp %s %s", cmd.PlayerId, target.Number)
}
