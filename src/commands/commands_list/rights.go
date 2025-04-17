package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/sirupsen/logrus"
	"strconv"
)

func SetRights(cmd *models.CommandsArgs) {
	logrus.Debugf("Set rights: %v", cmd.Params)
	if len(cmd.Params) < 2 {
		cmd.RconText(cmd.Usage)
		return
	}

	level, err := strconv.Atoi(cmd.Params[1])
	if err != nil || level < 0 || level > 100 {
		cmd.RconText("Please enter a number [0-100]. %s is not valid.", cmd.Params[1])
	} else {
		player, err := cmd.Context.Players.GetPlayer(cmd.Params[0])
		logrus.Debugf("player: %v", player)
		if err != nil {
			cmd.RconText(err.Error())
			return
		}
		cmd.Context.Players.UpdatePlayerRights(player.Number, level)
	}
}
