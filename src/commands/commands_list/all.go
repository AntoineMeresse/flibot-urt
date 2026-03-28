package commandslist

import (
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func All(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconUsage()
		return
	}

	player, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	if err != nil {
		cmd.RconText("^1Could not find your player info.")
		return
	}

	message := "[" + cmd.Context.UrtConfig.ApiConfig.ServerName + "] " + strings.Join(cmd.Params, " ")
	if err := cmd.Context.Api.SendGlobalMessage(player.Name, message); err != nil {
		cmd.RconText("^1Could not send global message: %s", err.Error())
	}
}
