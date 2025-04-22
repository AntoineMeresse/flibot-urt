package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
)

func CurrentMap(cmd *context.CommandsArgs) {
	cmd.RconText(cmd.Context.GetCurrentMap())
}
