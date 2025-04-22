package commandslist

import (
	"fmt"
	"github.com/AntoineMeresse/flibot-urt/src/context"
	"time"

	"github.com/AntoineMeresse/flibot-urt/src/api"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/AntoineMeresse/flibot-urt/src/utils/msg"
)

func MapGet(cmd *context.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconText("Please specify one or more maps.")
	} else {
		for _, m := range utils.CleanDuplicateElements(cmd.Params) {
			go downloadMap(m, cmd)
		}
	}
}

func downloadMap(mapSearch string, cmd *context.CommandsArgs) {
	// context.SetMapList()
	unique, mapname := uniqueMapExist(mapSearch, cmd)
	if unique {
		if !cmd.Context.IsMapAlreadyDownloaded(mapname) {
			newFile := fmt.Sprintf("%s/%s.pk3", cmd.Context.UrtConfig.DownloadPath, mapname)
			url := fmt.Sprintf("%s/%s", cmd.Context.UrtConfig.MapRepository, mapname)
			cmd.RconText(msg.DOWNLOAD_START, mapname)
			start := time.Now()
			if bytes, err := api.DownloadFile(newFile, url); err == nil {
				elapsed := time.Since(start)
				cmd.RconText(msg.DOWNLOAD_OK, mapname, utils.BytesNumberConverter(bytes), elapsed.Round(time.Millisecond))
				cmd.Context.MapSync()
			} else {
				cmd.RconText(msg.DOWNLOAD_KO, mapname)
			}
		} else {
			cmd.RconText(msg.DOWNLOAD_ALREADY_ON_SERV, mapname)
		}
	}
}

func uniqueMapExist(search string, cmd *context.CommandsArgs) (bool, string) {
	maps := cmd.Context.Api.GetMapsWithPattern(search)
	if len(maps) == 0 {
		cmd.RconText(msg.DOWNLOAD_NO_MAP, search)
		return false, ""
	} else if len(maps) > 1 {
		cmd.RconText(msg.DOWNLOAD_MULTIPLE_MAP, len(maps), search)
		return false, ""
	}
	return true, maps[0]
}
