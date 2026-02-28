package commandslist

import (
	"fmt"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"

	"github.com/AntoineMeresse/flibot-urt/src/vote"
)

func Callvote(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		voteList := []string{"Votes: "}
		for key, value := range vote.Votes {
			voteList = append(voteList, fmt.Sprintf("---> %s: ^5!callvote ^6%s", key, value.Usage))
		}
		cmd.RconList(voteList)
	} else {
		params := cmd.Params
		if (params[0] == "map" || params[0] == "nextmap") && len(params) > 1 {
			indexStr := ""
			if len(params) > 2 {
				indexStr = params[2]
			}
			mapName, ok := resolveMap(cmd, params[1], indexStr)
			if !ok {
				return
			}
			params = []string{params[0], mapName}
		}
		v := models.Vote{Params: params, PlayerId: cmd.PlayerId}
		cmd.Context.NewVote(v)
	}
}
