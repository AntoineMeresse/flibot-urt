package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Slap(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) > 0 {
		player, err := cmd.Context.Players.GetPlayer(cmd.Params[0])
		if err == nil {
			cmd.RconCommand("slap %s", player.Number)
			cmd.RconGlobalText("^5%s^7 was slapped!", player.Name)
		} else {
			cmd.RconText("%s", err.Error())
		}
	} else {
		cmd.RconUsage()
	}
}
