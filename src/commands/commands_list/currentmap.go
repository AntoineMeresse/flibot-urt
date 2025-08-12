package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func CurrentMap(cmd *appcontext.CommandsArgs) {
	cmd.RconText(cmd.Context.GetCurrentMap())
}
