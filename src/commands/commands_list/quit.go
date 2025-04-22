package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
)

func Quit(cmd *context.CommandsArgs) {
	cmd.RconCommand("kick %s \"%s\"", cmd.PlayerId, "This player isn't good enough for this map !")
}
