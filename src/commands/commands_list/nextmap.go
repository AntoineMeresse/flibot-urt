package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func NextMap(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconText(cmd.Context.GetNextMap())
		return
	}
	ChangeNextMap(cmd)
}
