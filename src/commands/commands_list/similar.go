package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	log "github.com/sirupsen/logrus"
)

func Similar(cmd *appcontext.CommandsArgs) {
	mapName := cmd.Context.GetCurrentMap()
	if len(cmd.Params) == 1 {
		mapName = cmd.Params[0]
	}

	maps, err := cmd.Context.Api.GetSimilarMaps(mapName)
	if err != nil {
		log.Errorf("[Similar] Error while trying to get infos from Api: %s", err.Error())
		cmd.RconText("Could not find similar maps for (%s)", mapName)
		return
	}

	m := "maps"
	if len(maps) == 1 {
		m = "map"
	}
	cmd.RconText("^5%d^7 random %s similar to %s:", len(maps), m, mapName)
	for _, line := range utils.ToShorterChunkArray(utils.AlternateColors(maps)) {
		cmd.RconText("  ^6|--->^7 %s", line)
	}
}
