package commandslist

import (
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"

	log "github.com/sirupsen/logrus"
)

func MapInformation(cmd *appcontext.CommandsArgs) {
	mapName := cmd.Context.GetCurrentMap()
	if len(cmd.Params) == 1 {
		mapName = cmd.Params[0]
	}

	infos, err := cmd.Context.Api.GetMapInformation(mapName)

	if err != nil {
		log.Errorf("[MapInformation] Error while trying to get infos from Api: %s", err.Error())
		cmd.RconText("Could not find map information for (%s)", mapName)
		return
	}

	cmd.RconText("^7Map infos for : ^5%s^7", infos.Mapname)
	cmd.RconText("^7 |--------> ^8Number of jumps^7 : %s", infos.Jumps)
	cmd.RconText("^7 |--------> ^8Level^7 : %d", infos.Level)
	cmd.RconText("^7 |--------> ^8Release Date^7 : %s", infos.Release)
	if len(infos.Types) > 0 {
		cmd.RconText("^7 |--------> ^8Types^7 : %s", strings.Join(infos.Types, " | "))
	}
	if len(infos.Notes) > 0 {
		cmd.RconText("^7 |--------> ^8Notes^7 : %s", strings.Join(infos.Notes, " | "))
	}
}
