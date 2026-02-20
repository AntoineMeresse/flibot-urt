package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Password(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconCommand("g_password \"\"")
		cmd.RconText("^7Server password ^2removed^7.")
	} else {
		password := cmd.Params[0]
		cmd.RconCommand("g_password \"%s\"", password)
		cmd.RconText("^7Server password set to: ^5%s^7.", password)
	}
}
