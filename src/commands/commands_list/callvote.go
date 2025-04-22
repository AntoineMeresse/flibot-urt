package commandslist

import (
	"fmt"
	"github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"

	"github.com/AntoineMeresse/flibot-urt/src/vote"
)

func Callvote(cmd *context.CommandsArgs) {
	if len(cmd.Params) == 0 {
		voteList := []string{"Votes: "}
		for key, value := range vote.Votes {
			voteList = append(voteList, fmt.Sprintf("---> %s: ^5!callvote ^6%s", key, value.Usage))
		}
		cmd.RconList(voteList)
	} else {
		v := models.Vote{Params: cmd.Params, PlayerId: cmd.PlayerId}
		cmd.Context.NewVote(v)
		VoteYes(cmd)
	}
}
