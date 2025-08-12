package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Invisible(cmd *appcontext.CommandsArgs) {
	cmd.RconCommand("invisible %s", cmd.PlayerId)
}
