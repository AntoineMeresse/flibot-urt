package commandslist

import (
	"time"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/AntoineMeresse/flibot-urt/src/utils/msg"
)

func ChangeMap(cmd *appcontext.CommandsArgs) {
	args, force := utils.ExtractForceFlag(cmd.Params)

	if len(args) == 0 {
		cmd.RconUsage()
		return
	}

	mapName, err := cmd.Context.GetMapWithCriteria(args[0])
	if err != nil {
		cmd.RconText(err.Error())
		return
	}

	if n := cmd.Context.Runs.AnyRunning(); n > 0 && !force {
		cmd.RconText("^3%d^7 player(s) are currently running. Add ^3-f^7 to change map anyway.", n)
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
