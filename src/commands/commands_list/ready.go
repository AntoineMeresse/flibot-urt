package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Ready(cmd *appcontext.CommandsArgs) {
	cmd.RconCommand("ready %s", cmd.PlayerId)
}
