package commandslist

import (
	"strconv"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Unignore(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconUsage()
		return
	}

	callerGuid := cmd.GetPlayerGuid()
	if callerGuid == "" {
		cmd.RconText("^1Could not identify your player.")
		return
	}

	id, err := strconv.Atoi(cmd.Params[0])
	if err != nil {
		cmd.RconText("^1Invalid id: %s", cmd.Params[0])
		return
	}

	target, ok := cmd.Context.DB.GetPlayerById(id)
	if !ok {
		cmd.RconText("^1No player found with id ^3%d^1.", id)
		return
	}

	if err := cmd.Context.DB.RemoveIgnore(callerGuid, target.Guid); err != nil {
		cmd.RconText("^1Error removing ignore: %s", err.Error())
		return
	}

	cmd.RconText("^7%s^7 removed from your ignore list.", target.Name)
}
