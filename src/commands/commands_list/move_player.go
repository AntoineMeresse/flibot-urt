package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func MovePlayer(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) < 2 {
		cmd.RconUsage()
		return
	}

	p1, err := cmd.Context.Players.GetPlayer(cmd.Params[0])
	if err != nil {
		cmd.RconText("%s", err.Error())
		return
	}

	p2, err := cmd.Context.Players.GetPlayer(cmd.Params[1])
	if err != nil {
		cmd.RconText("%s", err.Error())
		return
	}

	if p1.Number == p2.Number {
		cmd.RconText("^7Both players are the same.")
		return
	}

	cmd.RconCommand("forceteam %s free", p1.Number)
	cmd.RconCommand("tp %s %s", p1.Number, p2.Number)
}
