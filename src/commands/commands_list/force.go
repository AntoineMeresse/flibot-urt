package commandslist

import (
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

var teamAliases = map[string]string{
	"red":       "red",
	"r":         "red",
	"re":        "red",
	"blue":      "blue",
	"b":         "blue",
	"bl":        "blue",
	"blu":       "blue",
	"spec":      "spectator",
	"spectator": "spectator",
	"s":         "spectator",
	"sp":        "spectator",
	"spe":       "spectator",
	"green":     "green",
	"free":      "free",
}

func Force(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) < 2 {
		cmd.RconUsage()
		return
	}

	target, err := cmd.Context.Players.GetPlayer(cmd.Params[0])
	if err != nil {
		cmd.RconText("%s", err.Error())
		return
	}

	team, ok := teamAliases[strings.ToLower(cmd.Params[1])]
	if !ok {
		cmd.RconText("^1Invalid team ^3%s^1. Use: red, blue, spectator, green, free.", cmd.Params[1])
		return
	}

	cmd.RconCommand("forceteam %s %s", target.Number, team)
}
