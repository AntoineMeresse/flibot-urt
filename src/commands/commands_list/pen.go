package commandslist

import (
	"time"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Pen(cmd *appcontext.CommandsArgs) {
	player, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	if err != nil {
		cmd.RconText(err.Error())
		return
	}

	forceReroll := len(cmd.Params) > 0 && cmd.Params[0] == "-f"

	dayOfYear := time.Now().YearDay()
	yearlyUsed, err := cmd.Context.DB.PenGetYearlyAttempts(player.Guid)
	if err != nil {
		cmd.RconText(err.Error())
		return
	}
	remaining := dayOfYear - yearlyUsed

	if !forceReroll {
		currentSize, todayErr := cmd.Context.DB.PenPlayerGetDailySize(player.Guid)
		if todayErr == nil {
			// Already rolled today, just show current size
			cmd.RconGlobalText("^5B===D^7 %s pen(!s) size : ^5%.3f^7 cm", player.Name, currentSize)
			penGambleHint(cmd, remaining)
			return
		}
	}

	if remaining <= 0 {
		cmd.RconText("^7No attempts left! Your luck has run dry for today. Come back tomorrow! ^5:)")
		return
	}

	size := utils.RandomFloat(0., 50., 5)
	if err = cmd.Context.DB.PenAdd(player.Guid, size); err != nil {
		cmd.RconText(err.Error())
		return
	}

	remaining--
	cmd.RconGlobalText("^5B===D^7 %s pen(!s) size : ^5%.3f^7 cm", player.Name, size)
	penGambleHint(cmd, remaining)
}

func penGambleHint(cmd *appcontext.CommandsArgs, remaining int) {
	if remaining > 0 {
		cmd.RconText("^7Not satisfied? You've got ^5%d^7 attempt(s) left to gamble --> ^6!pen -f", remaining)
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

	cmd.RconText("^7=========== ^2Pen Hall Of Fame (^5%d^2) ^7===========", time.Now().Year())
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

	cmd.RconText("^7=========== ^1Pen Hall Of Shame (^5%d^2) ^7===========", time.Now().Year())
	for _, data := range datas {
		cmd.RconText("B===D %s - %.3f - %s", data.GetName(), data.Size, data.GetDate())
	}
}
