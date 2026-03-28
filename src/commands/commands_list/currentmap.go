package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func CurrentMap(cmd *appcontext.CommandsArgs) {
	cmd.RconText("^7[^5%s^7] ^3%s", cmd.Context.UrtConfig.ApiConfig.ServerName, cmd.Context.GetCurrentMap())
}
