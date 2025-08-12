package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Spec(cmd *appcontext.CommandsArgs) {
	cmd.RconCommand("forceteam %s spec", cmd.PlayerId)
}
