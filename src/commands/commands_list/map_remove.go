package commandslist

import (
	"fmt"
	"os"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func MapRemove(cmd *models.CommandsArgs) {
	mapSearch := cmd.Context.GetCurrentMap()
	if len(cmd.Params) > 0 {
		mapSearch = cmd.Params[0]
	}

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

