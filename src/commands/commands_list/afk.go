package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Afk(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) > 0 {
		player, err := cmd.Context.Players.GetPlayer(cmd.Params[0])
		if err == nil {
			if cmd.Context.Runs.IsRunning(player.Number) {
				cmd.RconText("^5%s^3 is currently running, can't move to spec.", player.Name)
				return
			}
			cmd.RconCommand("forceteam %s spec", player.Number)
		} else {
			cmd.RconText(err.Error())
		}
	}
}
