package commandslist

import (
	"fmt"
	"time"

	"github.com/AntoineMeresse/flibot-urt/src/api"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Download(server *models.Server, playerNumber string, params []string, isGlobal bool) {
	if len(params) == 0 {
		server.RconText("Please specify one or more maps.", isGlobal, playerNumber);
	} else {
		for _, m := range utils.CleanDuplicateElements(params) {
			go downloadMap(m, server, isGlobal, playerNumber)
		}
	}
}

func downloadMap(mapSearch string, server *models.Server, isGlobal bool, playerNumber string) {
	// server.SetMapList()
	unique, mapname := uniqueMapExist(mapSearch, server, isGlobal, playerNumber)
	if unique {
		if !server.IsMapAlreadyDownloaded(mapname) {
			newFile := fmt.Sprintf("%s/%s.pk3", server.UrtConfig.DownloadPath, mapname)
			url := fmt.Sprintf("%s/%s", server.UrtConfig.MapRepository, mapname)
			server.RconText(fmt.Sprintf("Downloading map %s: Start", mapname), isGlobal, playerNumber)
			start := time.Now()
			if err := api.DownloadFile(newFile, url); err == nil {
				elapsed := time.Since(start)
				server.RconText(fmt.Sprintf("Downloading map %s: OK (%s)", mapname, elapsed), isGlobal, playerNumber)
				server.SetMapList()
			} else {
				server.RconText(fmt.Sprintf("Downloading map %s: KO", mapname), isGlobal, playerNumber)
			}
		} else {
			server.RconText(fmt.Sprintf("%s is already on server !", mapname), isGlobal, playerNumber)
		}
	}
}

func uniqueMapExist(search string, server *models.Server, isGlobal bool, playerNumber string) (bool, string) {
	maps := server.Api.GetMapsWithPattern(search)
	if len(maps) == 0 {
		server.RconText(fmt.Sprintf("No map was found matching (%s)", search), isGlobal, playerNumber)
		return false, ""
	} else if len(maps) > 1 {
		server.RconText(fmt.Sprintf("Multiple maps found [%d] matching (%s)", len(maps), search), isGlobal, playerNumber)
		return false, ""
	}
	return true, maps[0]
}