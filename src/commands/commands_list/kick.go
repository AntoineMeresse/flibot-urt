package commandslist

import (
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Kick(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconUsage()
		return
	}

	target, ok := cmd.ResolveAdminTarget(cmd.Params[0])
	if !ok {
		return
	}

	reason := "Kicked by admin."
	if len(cmd.Params) > 1 {
		reason = strings.Join(cmd.Params[1:], " ")
	}

	cmd.RconCommand("kick %s \"%s\"", target.Number, reason)
	cmd.RconText("^5%s^7 was kicked! (^3%s^7)", target.Name, reason)
}
