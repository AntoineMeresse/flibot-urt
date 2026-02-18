package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/sirupsen/logrus"
)

const registeredRole = 20

func Register(cmd *appcontext.CommandsArgs) {
	player, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	if err != nil {
		cmd.RconText(err.Error())
		return
	}

	if err := cmd.Context.DB.SetPlayerRole(player.Guid, registeredRole); err != nil {
		logrus.Errorf("[Register] Error: %v", err)
		cmd.RconText("An error occurred while performing your registration. Please try again and contact an admin if the problem persists.")
		return
	}

	player.Role = registeredRole
	cmd.RconText("Successfully registered! Player id: %s", player.Id)
}
