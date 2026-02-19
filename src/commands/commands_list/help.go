package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Help(cmd *appcontext.CommandsArgs) {
	cmd.RconList(cmd.Params)
}
