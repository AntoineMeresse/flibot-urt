package commandslist

import "github.com/AntoineMeresse/flibot-urt/src/models"

func VoteYes(cmd *models.CommandsArgs) {
	v := models.Vote{Params: []string{"+"}, PlayerId: cmd.PlayerId}
	cmd.Server.NewVote(v)
}