package actionslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/commands"
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Say(action_params []string, server *models.Server) {
	if len(action_params) >= 3 {
		commands.HandleCommand(action_params, server)
	}
}