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

	mapName := utils.RandomValueFromSlice(maps)
	cmd.RconText("^7Next map rolled: ^5%s", mapName)
	cmd.Context.SetNextMap(mapName)
}
