package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func VoteYes(cmd *context.CommandsArgs) {
	v := models.Vote{Params: []string{"+"}, PlayerId: cmd.PlayerId}
	cmd.Context.NewVote(v)
}
