package commandslist

import (
	"fmt"
	"os"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func MapRemove(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		removeMap(cmd, cmd.Context.GetCurrentMap())
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

	removeMap(cmd, mapName)
}

func removeMap(cmd *appcontext.CommandsArgs, mapName string) {
	path := fmt.Sprintf("%s/%s.pk3", cmd.Context.UrtConfig.DownloadPath, mapName)
	if os.Remove(path) != nil {
		cmd.RconText("Error while trying to remove: ^5%s", mapName)
		return
	}

	cmd.RconText("^7Map (^5%s^7) has been removed.", mapName)
	cmd.Context.MapSync()
}
