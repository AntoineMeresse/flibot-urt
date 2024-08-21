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

func PenOfTheDay(cmd *models.CommandsArgs) {
	date, datas, err := cmd.Context.DB.Pen_PenOfTheDay()

	if err != nil {
		cmd.RconText(err.Error())
		return
	} 

	cmd.RconText("^7=========== ^6Pen of the day ^7(^5%s^7) ===========", date)
	for _, data := range(datas) {
		cmd.RconText("Pen ===> %s - %.3f ", data.GetName(), data.Size)
	}
}

func PenHallOfFame(cmd *models.CommandsArgs) {
	datas, err := cmd.Context.DB.Pen_PenHallOfFame()

	if err != nil {
		cmd.RconText(err.Error())
		return
	} 

	cmd.RconText("^7=========== ^2Pen Hall Of Fame ^7===========")
	for _, data := range(datas) {
		cmd.RconText("Pen ===> %s - %.3f - %s", data.GetName(), data.Size, data.GetDate())
	}
}

func PenHallOfShame(cmd *models.CommandsArgs) {
	datas, err := cmd.Context.DB.Pen_PenHallOfShame()

	if err != nil {
		cmd.RconText(err.Error())
		return
	} 

	cmd.RconText("^7=========== ^1Pen Hall Of Shame ^7===========")
	for _, data := range(datas) {
		cmd.RconText("Pen ===> %s - %.3f - %s", data.GetName(), data.Size, data.GetDate())
	}
}