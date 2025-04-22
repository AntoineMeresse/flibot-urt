package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
)

func Stamina(cmd *context.CommandsArgs) {
	cmd.RconCommand("customstamina %s", cmd.PlayerId)
}
