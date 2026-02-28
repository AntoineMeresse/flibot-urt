package commandslist

import (
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Kick(cmd *appcontext.CommandsArgs) {
	args, force := utils.ExtractForceFlag(cmd.Params)
	if len(args) == 0 {
		cmd.RconUsage()
		return
	}

	target, ok := cmd.ResolveAdminTarget(args[0])
	if !ok {
		return
	}

	if cmd.Context.Runs.IsRunning(target.Number) && !force {
		cmd.RconText("^5%s^3 is currently running. Add ^3-f^7 to kick anyway.", target.Name)
		return
	}

	reason := "Kicked by admin."
	if len(args) > 1 {
		reason = strings.Join(args[1:], " ")
	}

	cmd.RconCommand("kick %s \"%s\"", target.Number, reason)
	cmd.RconText("^5%s^7 was kicked! (^3%s^7)", target.Name, reason)
}
