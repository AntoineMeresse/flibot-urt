package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func PersonalBest(cmd *appcontext.CommandsArgs) {
	mapName := cmd.Context.GetCurrentMap()
	if len(cmd.Params) == 1 {
		mapName = cmd.Params[0]
	}

	player, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	if err != nil {
		cmd.RconText(err.Error())
		return
	}

	cmd.Context.SendPersonalBest(cmd.PlayerId, mapName, player.Guid, cmd.IsGlobal)
}
