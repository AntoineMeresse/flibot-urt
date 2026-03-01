package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func SetNextMap(cmd *appcontext.CommandsArgs) {
	ChangeNextMap(cmd)
}

func ChangeNextMap(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconUsage()
		return
	}

	indexStr := ""
	if len(cmd.Params) > 1 {
		indexStr = cmd.Params[1]
	}

	mapName, ok := resolveMap(cmd, cmd.Params[0], indexStr)
	if !ok {
		return
	}

	if mapName == cmd.Context.GetCurrentMap() && len(cmd.Context.GetMapList()) > 1 {
		cmd.RconText("^7Can't set nextmap to the current map.")
		return
	}

	cmd.RconText("Changing nextmap to: ^5%s", mapName)
	cmd.Context.SetNextMap(mapName)
}
