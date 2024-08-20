package commandslist

import "github.com/AntoineMeresse/flibot-urt/src/models"

func VoteVeto(cmd *models.CommandsArgs) {
	cmd.Context.RconCommand("veto")
	
	v := models.Vote{Params: []string{"v"}, PlayerId: cmd.PlayerId}
	cmd.Context.NewVote(v)
}