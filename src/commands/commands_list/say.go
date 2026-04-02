package commandslist

import (
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Say(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconUsage()
		return
	}

	player, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	if err != nil {
		cmd.RconText("^1Could not find your player info.")
		return
	}

	message := strings.Join(cmd.Params, " ")
	cmd.Context.RconText(true, "", "^7%s^7: %s", player.Name, message)
}
