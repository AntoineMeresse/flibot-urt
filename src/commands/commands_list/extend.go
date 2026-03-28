package commandslist

import (
	"fmt"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Extend(cmd *appcontext.CommandsArgs) {
	extendTime := 60

	if len(cmd.Params) > 0 {
		t, err := utils.ExtractNumber(cmd.Params[0])
		if err != nil || t <= 0 || t >= 1000 {
			cmd.RconUsage()
			return
		}
		extendTime = t
	}

	player, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	if err != nil {
		cmd.RconText("^1Could not identify your player.")
		return
	}

	if player.Role < 70 {
		v := models.Vote{Params: []string{"extend", fmt.Sprintf("%d", extendTime)}, PlayerId: cmd.PlayerId}
		cmd.Context.NewVote(v)
	} else {
		cmd.RconCommand("extend %d", extendTime)
	}
}
