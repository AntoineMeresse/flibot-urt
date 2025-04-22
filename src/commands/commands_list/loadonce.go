package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
)

func Loadonce(cmd *context.CommandsArgs) {
	cmd.RconCommand("simpleload %s", cmd.PlayerId)
}
