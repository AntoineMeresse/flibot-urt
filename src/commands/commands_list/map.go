package commandslist

import (
	"time"

	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils/msg"
)

func MapFn(cmd *models.CommandsArgs) {
	if len(cmd.Params) == 1 {
		mapName, err := cmd.Context.GetMapWithCriteria(cmd.Params[0])
		if err != nil {
			cmd.RconText(err.Error())
		} else {
			cmd.RconBigText(msg.MAP_CHANGE, *mapName)
			time.Sleep(200 * time.Millisecond)
			cmd.RconCommand("map %s", *mapName)
			cmd.Context.SetMapName(*mapName)
		}
	} else {
		cmd.RconUsage(cmd.Usage)
	}
}