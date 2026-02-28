package commandslist

import (
	"time"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils/msg"
)

func MapFn(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) != 1 {
		cmd.RconUsage()
		return
	}

	mapName, err := cmd.Context.GetMapWithCriteria(cmd.Params[0])
	if err != nil {
		cmd.RconText(err.Error())
		return
	}

	player, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	if err != nil {
		cmd.RconText("^1Could not identify your player.")
		return
	}

	if player.Role <= 60 {
		v := models.Vote{Params: []string{"map", *mapName}, PlayerId: cmd.PlayerId}
		cmd.Context.NewVote(v)
	} else {
		cmd.RconBigText(msg.MAP_CHANGE, *mapName)
		time.Sleep(200 * time.Millisecond)
		cmd.RconCommand("map %s", *mapName)
		cmd.Context.SetMapName(*mapName)
	}
}
