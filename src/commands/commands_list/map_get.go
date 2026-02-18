package commandslist

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"

	"github.com/AntoineMeresse/flibot-urt/src/api"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/AntoineMeresse/flibot-urt/src/utils/msg"
)

func MapGet(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconText("Please specify one or more maps.")
		return
	}

	maps := utils.CleanDuplicateElements(cmd.Params)
	var wg sync.WaitGroup
	var downloaded atomic.Bool

	for _, m := range maps {
		wg.Add(1)
		go downloadMapWorker(m, cmd, &wg, &downloaded)
	}

	wg.Wait()
	if downloaded.Load() {
		cmd.Context.MapSync()
	}
}

func downloadMapWorker(mapSearch string, cmd *appcontext.CommandsArgs, wg *sync.WaitGroup, downloaded *atomic.Bool) {
	defer wg.Done()
	if downloadMap(mapSearch, cmd) {
		downloaded.Store(true)
	}
}

func downloadMap(mapSearch string, cmd *appcontext.CommandsArgs) bool {
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
				return true
			}
			cmd.RconText(msg.DOWNLOAD_KO, mapname)
		} else {
			cmd.RconText(msg.DOWNLOAD_ALREADY_ON_SERV, mapname)
		}
	}
	return false
}

func uniqueMapExist(search string, cmd *appcontext.CommandsArgs) (bool, string) {
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
