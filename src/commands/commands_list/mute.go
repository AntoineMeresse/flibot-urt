package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Mute(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconUsage()
		return
	}

	target, ok := cmd.ResolveAdminTarget(cmd.Params[0])
	if !ok {
		return
	}

	cmd.RconCommand("mute %s", target.Number)
	cmd.RconText("^5%s^7 was muted!", target.Name)
}
