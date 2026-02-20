package commandslist

import (
	"strings"
	"time"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func penAsciiArt(size float64) string {
	return "8" + strings.Repeat("=", int(size/5)+1) + "D"
}

func penSizeCategory(size float64) string {
	switch {
	case size < 5:
		return "^7Don't try to improve your ^5micro^7 pen"
	case size < 10:
		return "^7Don't try to improve your ^5little^7 pen"
	case size < 18:
		return "^7Don't try to improve your ^5average^7 pen"
	case size < 24:
		return "^7Don't try to improve your ^5big^7 pen"
	default:
		return "^7Don't try to improve your ^5MONSTER^7 pen"
	}
}

func penRankDisplay(rank int, hos bool) (string, string) {
	switch rank {
	case 1:
		if hos {
			return "^3", "8=D"
		}
		return "^3", "8====D"
	case 2:
		if hos {
			return "^2", "8==D"
		}
		return "^2", "8===D"
	case 3:
		if hos {
			return "^8", "8===D"
		}
		return "^8", "8==D"
	default:
		if hos {
			return "^7", "8====D"
		}
		return "^7", "8=D"
	}
}

func Pen(cmd *appcontext.CommandsArgs) {
	player, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	if err != nil {
		cmd.RconText(err.Error())
		return
	}

	forceReroll := len(cmd.Params) > 0 && cmd.Params[0] == "-f"

	dayOfYear := time.Now().YearDay()
	attempts, err := cmd.Context.DB.PenGetAttempts(player.Guid)
	if err != nil {
		cmd.RconText(err.Error())
		return
	}
	remaining := dayOfYear - attempts

	if !forceReroll {
		currentSize, todayErr := cmd.Context.DB.PenPlayerGetDailySize(player.Guid)
		if todayErr == nil {
			cmd.RconGlobalText("%s, try again tomorrow !", penSizeCategory(currentSize))
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

	if err = cmd.Context.DB.PenIncrementAttempts(player.Guid); err != nil {
		cmd.RconText(err.Error())
		return
	}

	remaining--
	cmd.RconGlobalText("^5%s^7 %s pen(!s) size : ^5%.3f^7 cm", penAsciiArt(size), player.Name, size)
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
	if len(datas) == 0 {
		cmd.RconText("^7Use ^5!pen^7, there is no pen values yet ^1:(")
		return
	}
	for i, data := range datas {
		color, pen := penRankDisplay(i+1, false)
		cmd.RconText("%s%s ^7%s : ^5%.3f^7 cm.", color, pen, data.GetName(), data.Size)
	}
}

func PenHallOfFame(cmd *appcontext.CommandsArgs) {
	datas, err := cmd.Context.DB.PenPenHallOfFame()
	if err != nil {
		cmd.RconText(err.Error())
		return
	}

	cmd.RconText("^7=========== ^2Pen Hall Of Fame (^5%d^2) ^7===========", time.Now().Year())
	if len(datas) == 0 {
		cmd.RconText("^7Use ^5!pen^7, there is no pen values yet ^1:(")
		return
	}
	for i, data := range datas {
		color, pen := penRankDisplay(i+1, false)
		cmd.RconText("%s%s ^7%s : ^5%.3f^7 cm. (%s)", color, pen, data.GetName(), data.Size, data.GetDate())
	}
}

func PenHallOfShame(cmd *appcontext.CommandsArgs) {
	datas, err := cmd.Context.DB.PenPenHallOfShame()
	if err != nil {
		cmd.RconText(err.Error())
		return
	}

	cmd.RconText("^7=========== ^1Pen Hall Of Shame (^5%d^2) ^7===========", time.Now().Year())
	if len(datas) == 0 {
		cmd.RconText("^7Use ^5!pen^7, there is no pen values yet ^1:(")
		return
	}
	for i, data := range datas {
		color, pen := penRankDisplay(i+1, true)
		cmd.RconText("%s%s ^7%s : ^5%.3f^7 cm. (%s)", color, pen, data.GetName(), data.Size, data.GetDate())
	}
}
