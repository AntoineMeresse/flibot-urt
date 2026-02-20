package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func PenCoin(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) < 1 {
		cmd.RconText(cmd.Usage)
		return
	}

	player, err := cmd.Context.Players.GetPlayer(cmd.Params[0])
	if err != nil {
		cmd.RconText(err.Error())
		return
	}

	if !cmd.Context.GivePenCoin(*player) {
		cmd.RconText("^1Failed to give pencoin to ^5%s", player.Name)
		return
	}

	cmd.RconText("^5%s^7 just got a free extra pen attempt!", player.Name)
}
