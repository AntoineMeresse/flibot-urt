package commandslist

import (
	"fmt"
	"time"

	"github.com/AntoineMeresse/flibot-urt/src/api"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/AntoineMeresse/flibot-urt/src/utils/msg"
)

func Download(cmd *models.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconText("Please specify one or more maps.");
	} else {
		for _, m := range utils.CleanDuplicateElements(cmd.Params) {
			go downloadMap(m, cmd)
		}
	}
}

func downloadMap(mapSearch string, cmd *models.CommandsArgs) {
	// server.SetMapList()
	unique, mapname := uniqueMapExist(mapSearch, cmd)
	if unique {
		if !cmd.Server.IsMapAlreadyDownloaded(mapname) {
			newFile := fmt.Sprintf("%s/%s.pk3",cmd.Server.UrtConfig.DownloadPath, mapname)
			url := fmt.Sprintf("%s/%s", cmd.Server.UrtConfig.MapRepository, mapname)
			cmd.RconText(msg.DOWNLOAD_START, mapname)
			start := time.Now()
			if err := api.DownloadFile(newFile, url); err == nil {
				elapsed := time.Since(start)
				cmd.RconText(msg.DOWNLOAD_OK, mapname, elapsed)
				cmd.Server.SetMapList()
			} else {
				cmd.RconText(msg.DOWNLOAD_KO, mapname)
			}
		} else {
			cmd.RconText(msg.DOWNLOAD_ALREADY_ON_SERV, mapname)
		}
	}
}

func uniqueMapExist(search string, cmd *models.CommandsArgs) (bool, string) {
	maps := cmd.Server.Api.GetMapsWithPattern(search)
	if len(maps) == 0 {
		cmd.RconText(msg.DOWNLOAD_NO_MAP, search)
		return false, ""
	} else if len(maps) > 1 {
		cmd.RconText(msg.DOWNLOAD_MULTIPLE_MAP, len(maps), search)
		return false, ""
	}
	return true, maps[0]
}