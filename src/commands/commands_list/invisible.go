package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
)

func Invisible(cmd *context.CommandsArgs) {
	cmd.RconCommand("invisible %s", cmd.PlayerId)
}
