package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Kms(cmd *appcontext.CommandsArgs) {
	cmd.RconCommand("smite %s", cmd.PlayerId)
}
