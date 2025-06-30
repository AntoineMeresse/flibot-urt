package commandslist

import (
	"fmt"
	"strings"

	"github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/sirupsen/logrus"
)

func Register(cmd *context.CommandsArgs) {
	player, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	clientMsg := "Successfully registered!"
	if err == nil {
		errMsg := ""
		if id, errSavePlayer := cmd.Context.DB.SaveNewPlayer(player.Name, player.Guid, player.Ip); errSavePlayer != nil {
			logrus.Errorf("[Register] Error: %v", errSavePlayer)
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
