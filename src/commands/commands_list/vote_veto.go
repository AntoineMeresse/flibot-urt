package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func VoteVeto(cmd *appcontext.CommandsArgs) {
	cmd.Context.RconCommand("veto")

	v := models.Vote{Params: []string{"v"}, PlayerId: cmd.PlayerId}
	cmd.Context.NewVote(v)
}
