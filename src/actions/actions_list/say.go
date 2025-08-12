package actionslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/commands"
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Say(actionParams []string, c *appcontext.AppContext) {
	if len(actionParams) >= 3 {
		commands.HandleCommand(actionParams, c)
	}
}
