package actionslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/commands"
	"github.com/AntoineMeresse/flibot-urt/src/context"
)

func Say(actionParams []string, c *context.Context) {
	if len(actionParams) >= 3 {
		commands.HandleCommand(actionParams, c)
	}
}
