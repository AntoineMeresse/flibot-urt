package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	log "github.com/sirupsen/logrus"
)

func Missing(cmd *appcontext.CommandsArgs) {
	maxlvl := "100"
	if len(cmd.Params) == 1 {
		maxlvl = cmd.Params[0]
	}

	guid := cmd.GetPlayerGuid()
	if guid == "" {
		return
	}

	maps, err := cmd.Context.Api.GetMissingMaps(maxlvl, guid)
	if err != nil {
		log.Errorf("[Missing] Error while trying to get infos from Api: %s", err.Error())
		cmd.RconText("Could not find missing maps")
		return
	}

	m := "maps"
	if len(maps) == 1 {
		m = "map"
	}
	cmd.RconText("^5%d^7 random %s under lvl %s where you don't have a run:", len(maps), m, maxlvl)
	for _, line := range utils.ToShorterChunkArray(maps) {
		cmd.RconText("  ^6|--->^7 %s", line)
	}
}
