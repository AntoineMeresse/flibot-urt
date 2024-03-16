package commandslist

import (
	"time"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func MapFn(cmd *models.CommandsArgs) {
	if len(cmd.Params) == 1 {
		mapName, err := cmd.Server.GetMapWithCriteria(cmd.Params[0])
		if err != nil {
			cmd.RconText(err.Error())
		} else {
			cmd.RconBigText("^7Changing map to %s", *mapName)
			time.Sleep(200 * time.Millisecond)
			cmd.RconCommand("map %s", *mapName)
			cmd.Server.SetMapName(*mapName)
		}
	} else {
		cmd.RconUsage(cmd.Usage)
	}
}