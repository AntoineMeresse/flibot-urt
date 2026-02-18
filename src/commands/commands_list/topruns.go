package commandslist

import (
	"fmt"
	"log/slog"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func ToprunsInformation(cmd *appcontext.CommandsArgs) {
	displayRunsInformation(cmd, true)
}

func TopInformation(cmd *appcontext.CommandsArgs) {
	displayRunsInformation(cmd, false)
}

func displayRunsInformation(cmd *appcontext.CommandsArgs, displayAll bool) {
	mapName := cmd.Context.GetCurrentMap()
	if len(cmd.Params) == 1 {
		mapName = cmd.Params[0]
	}

	infos, err := cmd.Context.Api.GetToprunsInformation(mapName)

	if err != nil {
		slog.Error("ToprunsInformation: error from API", "err", err)
		cmd.RconText("Could not find topruns for (%s)", mapName)
		return
	}

	if displayAll {
		cmd.RconText("^7Topruns for :  ^5%s^7", infos.Mapname)
	} else {
		cmd.RconText("^7Toprun for :  ^5%s^7", infos.Mapname)
	}

	if len(infos.RunsInfos) == 0 {
		cmd.RconText("^7|--------> No runs found")
		return
	}

	waysNumber := 1
	slog.Debug("Start displayRunsInformation")
	for way, runinfos := range infos.RunsInfos {
		cmd.RconText("^7|-> Runs for ^5way %s^7 :", way)
		if !displayAll {
			runinfos = runinfos[:1]
		}
		for i, run := range runinfos {
			slog.Debug(fmt.Sprintf("Interation n°%d displayRunsInformation", i))
			cmd.RconText("^7|-------->%2d) %s%s^7 | %s | %s", i+1, utils.GetColorRun(i), run.RunTime, run.RunDate, run.PlayerName)
		}
		if waysNumber != len(infos.RunsInfos) {
			cmd.RconText("^7|")
		}
		waysNumber += 1
	}
	slog.Debug("End displayRunsInformation")
}
