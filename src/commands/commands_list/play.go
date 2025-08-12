package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Play(cmd *appcontext.CommandsArgs) {
	cmd.RconCommand("forceteam %s free", cmd.PlayerId)
}
