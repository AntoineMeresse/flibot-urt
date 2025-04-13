package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	log "github.com/sirupsen/logrus"
)

func PersonnalBest(cmd *models.CommandsArgs) {
	mapName := cmd.Context.Settings.Mapname
	if len(cmd.Params) == 1 {
		mapName = cmd.Params[0]
	}

	player, playerErr := cmd.Context.Players.GetPlayer(cmd.PlayerId)

	if playerErr != nil {
		cmd.RconText(playerErr.Error())
		return
	}

	infos, err := cmd.Context.Api.GetPersonalBestInformation(mapName, player.Guid)

	if err != nil {
		log.Errorf("[PersonnalBest] Error while trying to get infos from Api: %s", err.Error())
		cmd.RconText("Could not find pb for (%s)", mapName)
		return
	}

	cmd.RconText("^7|   ^3========^7 Personal best for:  ^5%s^7 ^3========^7", infos.Mapname)

	if len(infos.RunsInfos) == 0 {
		cmd.RconText("^7|--------> No runs found")
		return
	}

	waysNumber := 1
	for _, data := range infos.RunsInfos {
		cmd.RconText("^7|-> Runs for ^5way %s^7 :", data.Wayname)
		cmd.RconText("^7|-------->(^3%s^7) ^6%s^7 | %s | %s", data.Rank, data.Run.RunTime, data.Run.RunDate, utils.DecolorString(data.Run.PlayerName))
		if waysNumber != len(infos.RunsInfos) {
			cmd.RconText("^7|")
		}
		waysNumber += 1
	}
}
