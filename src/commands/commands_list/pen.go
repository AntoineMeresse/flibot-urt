package commandslist

import (
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/sirupsen/logrus"
)

func Pen(cmd *appcontext.CommandsArgs) {
	player, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	if err != nil {
		cmd.RconText(err.Error())
		return
	}

	size := utils.RandomFloat(0., 50., 5)
	err = cmd.Context.DB.PenAdd(player.Guid, size)

	pen := "B===D"

	if err != nil {
		logrus.Debugf("Error: %v | Type: %T", err, err)
		if strings.Contains(err.Error(), "duplicate key") {
			size, err := cmd.Context.DB.PenPlayerGetDailySize(player.Guid)
			if err != nil {
				return
			}
			cmd.RconText("You already use !pen today. Size: %.3f", size)
		}
	} else {
		cmd.RconGlobalText("^5%s^7 %s pen(!s) size : ^5%.3f^7 cm", pen, player.Name, size)
	}
}

func PenOfTheDay(cmd *appcontext.CommandsArgs) {
	date, datas, err := cmd.Context.DB.PenPenOfTheDay()

	if err != nil {
		cmd.RconText(err.Error())
		return
	}

	cmd.RconText("^7=========== ^6Pen of the day ^7(^5%s^7) ===========", date)
	for _, data := range datas {
		cmd.RconText("B===D %s - %.3f ", data.GetName(), data.Size)
	}
}

func PenHallOfFame(cmd *appcontext.CommandsArgs) {
	datas, err := cmd.Context.DB.PenPenHallOfFame()

	if err != nil {
		cmd.RconText(err.Error())
		return
	}

	cmd.RconText("^7=========== ^2Pen Hall Of Fame ^7===========")
	for _, data := range datas {
		cmd.RconText("B===D %s - %.3f - %s", data.GetName(), data.Size, data.GetDate())
	}
}

func PenHallOfShame(cmd *appcontext.CommandsArgs) {
	datas, err := cmd.Context.DB.PenPenHallOfShame()

	if err != nil {
		cmd.RconText(err.Error())
		return
	}

	cmd.RconText("^7=========== ^1Pen Hall Of Shame ^7===========")
	for _, data := range datas {
		cmd.RconText("B===D %s - %.3f - %s", data.GetName(), data.Size, data.GetDate())
	}
}
