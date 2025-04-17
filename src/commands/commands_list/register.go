package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"strings"
)

func Register(cmd *models.CommandsArgs) {
	player, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	clientMsg := "Successfully registered!"
	if err == nil {
		errMsg := ""
		if errSavePlayer := cmd.Context.DB.SaveNewPlayer(player.Name, player.Guid, player.Ip); errSavePlayer != nil {
			errMsg += "An error occurred while performing your registration."
		}
		if errInitRight := cmd.Context.DB.InitRight(player.Guid); errInitRight != nil {
			errMsg += " Could not initialize your rights."
		}
		if errMsg != "" {
			clientMsg = strings.Trim(errMsg, " ") + " Please try again and contact an admin if the problem persists."
		}
	}
	cmd.RconText(clientMsg)
}
