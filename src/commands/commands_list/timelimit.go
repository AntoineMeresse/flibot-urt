package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Timelimit(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 1 {
		t, err := utils.ExtractNumber(cmd.Params[0])

		if err == nil {
			if t > 0 && t < 1000 {
				cmd.RconCommand("setTimelimit %d", t)
				return
			}
		}
	}
	cmd.RconUsage()
}
