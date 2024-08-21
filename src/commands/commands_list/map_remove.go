package commandslist

import (
	"fmt"
	"os"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func MapRemove(cmd *models.CommandsArgs) {
	if len(cmd.Params) == 0 {
		removeMap(cmd, cmd.Context.GetCurrentMap())
	} else {
		for _, mapname := range(cmd.Params) {
			removeMap(cmd, mapname)
		}
	}
}

func removeMap(cmd *models.CommandsArgs, mapSearch string) {
	mapName, err := cmd.Context.GetMapWithCriteria(mapSearch)
	
	if err != nil {
		cmd.RconText(err.Error())
		return;
	} 

	path := fmt.Sprintf("%s/%s.pk3", cmd.Context.UrtConfig.DownloadPath, *mapName)
	if os.Remove(path) != nil {
		cmd.RconText("Error while trying to remove: %s", *mapName)
		return
	}

	
	cmd.RconText("^7Map (^5%s^7) has been removed.", *mapName)
	cmd.Context.MapSync()
}
