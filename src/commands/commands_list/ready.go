package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
)

func Ready(cmd *context.CommandsArgs) {
	cmd.RconCommand("ready %s", cmd.PlayerId)
}
