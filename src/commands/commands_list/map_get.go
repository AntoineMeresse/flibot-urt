package commandslist

import (
	"fmt"
	"strings"
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

	cmd.Context.SetMapList()

	maps := utils.CleanDuplicateElements(cmd.Params)
	var wg sync.WaitGroup
	var downloaded atomic.Bool

	for _, m := range maps {
		wg.Add(1)
		go func(mapSearch string) {
			defer wg.Done()
			if downloadMap(mapSearch, cmd) {
				downloaded.Store(true)
			}
		}(m)
	}

	wg.Wait()
	if downloaded.Load() {
		cmd.Context.MapSync()
	}
}

func downloadMap(mapSearch string, cmd *appcontext.CommandsArgs) bool {
	search := strings.TrimSuffix(mapSearch, ".pk3")

	maps := cmd.Context.Api.GetMapsWithPattern(search)
	if len(maps) == 0 {
		cmd.RconText(msg.DOWNLOAD_NO_MAP, search)
		return false
	}

	// Check for exact match
	for _, m := range maps {
		if strings.EqualFold(m, search) {
			maps = []string{m}
			break
		}
	}

	if len(maps) > 1 {
		handleMultipleMatches(cmd, search, maps)
		return false
	}

	return handleSingleMatch(cmd, maps[0])
}

func handleMultipleMatches(cmd *appcontext.CommandsArgs, search string, maps []string) {
	count := len(maps)
	displayMaps := maps
	if count > 6 {
		displayMaps = maps[:6]
	}

	cmd.RconText(msg.DOWNLOAD_MULTIPLE_MAP, count, search)
	for _, m := range displayMaps {
		if cmd.Context.IsMapAlreadyDownloaded(m) {
			cmd.RconText(msg.DOWNLOAD_MAP_ITEM_ALREADY, m)
		} else {
			cmd.RconText(msg.DOWNLOAD_MAP_ITEM, m)
		}
	}
	if count > 6 {
		cmd.RconText("^7|-> [...]")
	}
}

func handleSingleMatch(cmd *appcontext.CommandsArgs, mapname string) bool {
	if cmd.Context.IsMapAlreadyDownloaded(mapname) {
		cmd.RconText(msg.DOWNLOAD_ALREADY_ON_SERV, mapname)
		return false
	}

	newFile := fmt.Sprintf("%s/%s.pk3", cmd.Context.UrtConfig.DownloadPath, mapname)
	url := fmt.Sprintf("%s/%s.pk3", cmd.Context.UrtConfig.MapRepository, mapname)

	cmd.RconText(msg.DOWNLOAD_START, mapname)
	start := time.Now()

	bytes, err := api.DownloadFile(newFile, url)
	if err != nil {
		cmd.RconText(msg.DOWNLOAD_KO, mapname)
		return false
	}

	elapsed := time.Since(start)
	cmd.RconText(msg.DOWNLOAD_OK, mapname, utils.BytesNumberConverter(bytes), elapsed.Round(time.Millisecond))
	return true
}
