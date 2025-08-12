package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Loadonce(cmd *appcontext.CommandsArgs) {
	cmd.RconCommand("simpleload %s", cmd.PlayerId)
}
