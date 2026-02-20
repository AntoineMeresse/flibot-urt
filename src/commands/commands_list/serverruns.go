package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func ServerRuns(cmd *appcontext.CommandsArgs) {
	mapName := cmd.Context.GetCurrentMap()
	if len(cmd.Params) == 1 {
		mapName = cmd.Params[0]
	}

	guids := cmd.Context.Players.GetGuids()
	infos, err := cmd.Context.Api.GetServerRunsInformation(mapName, guids)
	if err != nil {
		log.Errorf("[ServerRuns] Error while trying to get infos from Api: %s", err.Error())
		cmd.RconText("Could not find server runs for (%s)", mapName)
		return
	}

	displayRunsInfos(cmd, infos, true, "Server runs")
}
