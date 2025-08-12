package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Stamina(cmd *appcontext.CommandsArgs) {
	cmd.RconCommand("customstamina %s", cmd.PlayerId)
}
