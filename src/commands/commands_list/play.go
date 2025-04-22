package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
)

func Play(cmd *context.CommandsArgs) {
	cmd.RconCommand("forceteam %s free", cmd.PlayerId)
}
