package commandslist

import "github.com/AntoineMeresse/flibot-urt/src/models"

func Callvote(cmd *models.CommandsArgs) {
	v := models.Vote{Params: cmd.Params, PlayerId: cmd.PlayerId}
	cmd.Context.NewVote(v)
}