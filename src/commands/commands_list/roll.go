package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Roll(cmd *appcontext.CommandsArgs) {
	maps := cmd.Context.GetMapList()
	if len(maps) == 0 {
		cmd.RconText("^7No maps available to roll.")
		return
	}

	currentMap := cmd.Context.GetCurrentMap()
	filtered := make([]string, 0, len(maps))
	for _, m := range maps {
		if m != currentMap {
			filtered = append(filtered, m)
		}
	}
	if len(filtered) == 0 {
		filtered = maps
	}

	mapName := utils.RandomValueFromSlice(filtered)
	cmd.RconText("^7Next map rolled: ^5%s", mapName)
	cmd.Context.SetNextMap(mapName)
}
