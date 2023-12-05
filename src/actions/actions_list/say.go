package actionslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/commands"
)

func Say(action_params []string) {
	if len(action_params) >= 3 {
		commands.HandleCommand(action_params)
	}
}