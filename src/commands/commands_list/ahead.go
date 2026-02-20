package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func Ahead(cmd *appcontext.CommandsArgs) {
	mapName := cmd.Context.GetCurrentMap()
	if len(cmd.Params) == 1 {
		mapName = cmd.Params[0]
	}

	guid := cmd.GetPlayerGuid()
	if guid == "" {
		return
	}

	infos, err := cmd.Context.Api.GetRunsAhead(mapName, guid)
	if err != nil {
		log.Errorf("[Ahead] Error while trying to get infos from Api: %s", err.Error())
		cmd.RconText("Could not find runs ahead for (%s)", mapName)
		return
	}

	displayRunsInfos(cmd, infos, true, "Runs ahead")
}
