package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func VoteNo(cmd *appcontext.CommandsArgs) {
	v := models.Vote{Params: []string{"-"}, PlayerId: cmd.PlayerId}
	cmd.Context.NewVote(v)
}
