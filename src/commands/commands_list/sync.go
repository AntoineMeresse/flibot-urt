package commandslist

import appcontext "github.com/AntoineMeresse/flibot-urt/src/context"

func Sync(cmd *appcontext.CommandsArgs) {
	cmd.Context.MapSync()
}
