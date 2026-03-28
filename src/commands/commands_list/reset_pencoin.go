package commandslist

import (
	"strconv"
	"time"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func ResetPenCoin(cmd *appcontext.CommandsArgs) {
	var amount int
	if len(cmd.Params) == 0 {
		amount = time.Now().YearDay() - 5
		if amount < 0 {
			amount = 0
		}
	} else {
		var err error
		amount, err = strconv.Atoi(cmd.Params[0])
		if err != nil || amount < 0 {
			cmd.RconText("^1Invalid amount: %s. Must be a positive number.", cmd.Params[0])
			return
		}
	}

	if err := cmd.Context.DB.PenSetAllAttempts(amount); err != nil {
		cmd.RconText("^1Failed to reset pen coins: %s", err.Error())
		return
	}

	cmd.RconText("^7Pen coins reset to ^5%d^7 for all players.", amount)
}
