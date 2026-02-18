package commandslist

import (
	"log/slog"
	"strconv"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func SetRights(cmd *appcontext.CommandsArgs) {
	slog.Debug("Set rights", "params", cmd.Params)
	if len(cmd.Params) < 2 {
		cmd.RconText(cmd.Usage)
		return
	}

	level, err := strconv.Atoi(cmd.Params[1])
	if err != nil || level < 0 || level > 100 {
		cmd.RconText("Please enter a number [0-100]. %s is not valid.", cmd.Params[1])
	} else {
		player, err := cmd.Context.Players.GetPlayer(cmd.Params[0])
		slog.Debug("Set rights player", "player", player)
		if err != nil {
			cmd.RconText(err.Error())
			return
		}
		cmd.Context.Players.UpdatePlayerRights(player.Number, level)
	}
}
