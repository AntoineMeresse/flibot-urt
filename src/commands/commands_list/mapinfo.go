package commandslist

import (
	"fmt"
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

	mapperLabel := "Mapper"
	if len(infos.Mappers) > 1 {
		mapperLabel = "Mappers"
	}

	cmd.RconText("^7Map info for : ^5%s^7 (%s)", infos.Mapname, infos.Filename)
	if len(infos.Mappers) > 0 {
		cmd.RconText("^7 |--------> ^8%s^7 : %s", mapperLabel, strings.Join(infos.Mappers, " | "))
	}
	cmd.RconText("^7 |--------> ^8Number of jumps^7 : %s", infos.Jumps)
	cmd.RconText("^7 |--------> ^8Level^7 : %d", infos.Level)
	cmd.RconText("^7 |--------> ^8Release Date^7 : %s", strings.Replace(infos.Release, " 00:00:00 GMT", "", 1))
	if len(infos.Types) > 0 {
		cmd.RconText("^7 |--------> ^8Types^7 : %s", strings.Join(infos.Types, " | "))
	}
	if len(infos.Notes) > 0 {
		cmd.RconText("^7 |--------> ^8Notes^7 : %s", strings.Join(infos.Notes, ", "))
	}
	if len(infos.Tags) > 0 {
		cmd.RconText("^7 |--------> ^8Tags^7 : %s", strings.Join(infos.Tags, " | "))
	}
	if len(infos.Functions) > 0 {
		cmd.RconText("^7 |--------> ^8Functions^7 : %s", strings.Join(infos.Functions, " | "))
	}
	var more []string
	if len(infos.Addons) > 0 {
		more = append(more, fmt.Sprintf("^5%d^7 mod(s)", len(infos.Addons)))
	}
	if len(infos.Videos) > 0 {
		more = append(more, fmt.Sprintf("^5%d^7 video(s)", len(infos.Videos)))
	}
	if len(more) > 0 {
		cmd.RconText("^7 |--------> ^8Available^7: %s", strings.Join(more, " | "))
	}
}
