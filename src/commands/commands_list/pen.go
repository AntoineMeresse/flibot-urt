package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Pen(cmd *models.CommandsArgs) {
	player, err  := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	if err != nil {
		cmd.RconText(err.Error())
		return;
	}

	size := utils.RandomFloat(0. , 50., 5)
	err = cmd.Context.DB.Pen_add(player.Guid, size)
	
	pen := "B===D"

	if err != nil {
		cmd.RconText(err.Error())
	} else {
		cmd.RconGlobalText("^5%s^7 %s pen(!s) size : ^5%.3f^7 cm", pen, player.Name, size)
	}
}