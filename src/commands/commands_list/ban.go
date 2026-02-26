package commandslist

import (
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Ban(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconUsage()
		return
	}

	target, ok := cmd.ResolveAdminTarget(cmd.Params[0])
	if !ok {
		return
	}

	reason := ""
	if len(cmd.Params) > 1 {
		reason = strings.Join(cmd.Params[1:], " ")
	}

	if err := cmd.Context.DB.AddBan(target.Guid, target.Ip, reason); err != nil {
		cmd.RconText("^1Error saving ban: %s", err.Error())
		return
	}

	cmd.RconCommand("kick %s \"Banned: %s\"", target.Number, reason)
	if reason != "" {
		cmd.RconText("^5%s^7 has been banned. (^3%s^7)", target.Name, reason)
	} else {
		cmd.RconText("^5%s^7 has been banned.", target.Name)
	}
}
