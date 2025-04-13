package actionslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/commands"
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Say(actionParams []string, context *models.Context) {
	if len(actionParams) >= 3 {
		commands.HandleCommand(actionParams, context)
	}
}
