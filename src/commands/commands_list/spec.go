package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
)

func Spec(cmd *context.CommandsArgs) {
	cmd.RconCommand("forceteam %s spec", cmd.PlayerId)
}
