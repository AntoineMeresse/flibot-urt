package commandslist

import (
	"sort"
	"time"

	"github.com/AntoineMeresse/flibot-urt/src/api"
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/maruel/natural"
	log "github.com/sirupsen/logrus"
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
		log.Errorf("[ToprunsInformation] Error while trying to get infos from Api: %s", err.Error())
		cmd.RconText("Could not find topruns for (%s)", mapName)
		return
	}

	label := "Topruns"
	if !displayAll {
		label = "Toprun"
	}
	displayRunsInfos(cmd, infos, displayAll, label)
}

func displayRunsInfos(cmd *appcontext.CommandsArgs, infos api.ToprunsInfos, displayAll bool, label string) {
	cmd.RconText("^7%s for :  ^5%s^7", label, infos.Mapname)

	if len(infos.RunsInfos) == 0 {
		cmd.RconText("^7|--------> No runs found")
		return
	}

	ways := make([]string, 0, len(infos.RunsInfos))
	for way := range infos.RunsInfos {
		ways = append(ways, way)
	}
	sort.Sort(natural.StringSlice(ways))

	total := 0
	for _, runs := range infos.RunsInfos {
		if !displayAll {
			total += 1
		} else {
			total += len(runs)
		}
	}
	needBreak := total > 50

	waysNumber := 1
	go func() {
		log.Debug("Start displayRunsInfos")
		for _, way := range ways {
			runinfos := infos.RunsInfos[way]
			cmd.RconText("^7|-> Runs for ^5way %s^7 :", way)
			if !displayAll {
				runinfos = runinfos[:1]
			}
			for i, run := range runinfos {
				log.Debugf("Iteration n°%d displayRunsInfos", i)
				cmd.RconText("^7|-------->%2d) %s%s^7 | %s | %s", i+1, utils.GetColorRun(i), run.RunTime, run.RunDate, run.PlayerName)
			}
			if waysNumber != len(ways) {
				cmd.RconText("^7|")
				if needBreak {
					time.Sleep(100 * time.Millisecond)
				}
			}
			waysNumber += 1
		}
		log.Debug("End displayRunsInfos")
	}()
}
