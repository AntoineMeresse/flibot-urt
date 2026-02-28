package commandslist

import (
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
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
	args, force := utils.ExtractForceFlag(cmd.Params)
	if len(args) < 2 {
		cmd.RconUsage()
		return
	}

	target, err := cmd.Context.Players.GetPlayer(args[0])
	if err != nil {
		cmd.RconText("%s", err.Error())
		return
	}

	team, ok := teamAliases[strings.ToLower(args[1])]
	if !ok {
		cmd.RconText("^1Invalid team ^3%s^1. Use: red, blue, spectator, green, free.", args[1])
		return
	}

	if cmd.Context.Runs.IsRunning(target.Number) && !force {
		cmd.RconText("^5%s^3 is currently running. Add ^3-f^7 to force anyway.", target.Name)
		return
	}

	cmd.RconCommand("forceteam %s %s", target.Number, team)
}
