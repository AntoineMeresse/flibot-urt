package commandslist

import (
	"fmt"
	"log/slog"
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Register(cmd *appcontext.CommandsArgs) {
	player, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	clientMsg := "Successfully registered!"
	if err == nil {
		errMsg := ""
		if id, errSavePlayer := cmd.Context.DB.SaveNewPlayer(player.Name, player.Guid, player.Ip); errSavePlayer != nil {
			slog.Error("Register error", "err", errSavePlayer)
			errMsg += "An error occurred while performing your registration."
		} else {
			clientMsg += fmt.Sprintf(" Player id: %d", id)
		}
		if errMsg != "" {
			clientMsg = strings.Trim(errMsg, " ") + " Please try again and contact an admin if the problem persists."
		}
	}
	cmd.RconText(clientMsg)
}
