package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
)

func NextMap(cmd *context.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconText(cmd.Context.GetNextMap())
		return
	}
	ChangeNextMap(cmd)
}
