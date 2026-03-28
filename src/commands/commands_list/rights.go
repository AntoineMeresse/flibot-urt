package commandslist

import (
	"strconv"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/sirupsen/logrus"
)

func SetRights(cmd *appcontext.CommandsArgs) {
	logrus.Debugf("Set rights: %v", cmd.Params)
	if len(cmd.Params) < 2 {
		cmd.RconText(cmd.Usage)
		return
	}

	level, err := strconv.Atoi(cmd.Params[1])
	if err != nil || level < -1 || level > 100 {
		cmd.RconText("Please enter a number [-1 to 100]. %s is not valid.", cmd.Params[1])
		return
	}

	player, ok := cmd.ResolveAdminTarget(cmd.Params[0])
	if !ok {
		return
	}
	logrus.Debugf("player: %v", player)

	caller, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	if err != nil {
		cmd.RconText("%s", err.Error())
		return
	}
	if level > caller.Role {
		cmd.RconText("^1You cannot assign a level higher than your own (%d)!", caller.Role)
		return
	}

	cmd.Context.Players.UpdatePlayerRights(player.Number, level)
	if err := cmd.Context.DB.SetPlayerRole(player.Guid, level); err != nil {
		logrus.Errorf("SetRights: failed to persist role for %s: %v", player.Guid, err)
	}
	cmd.RconText("^7%s^7 rights updated to ^5%d", player.Name, level)
}
