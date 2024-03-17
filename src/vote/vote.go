package vote

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/sirupsen/logrus"
)

func InitVoteSystem(voteChannel <-chan models.Vote, server *models.Server) {
	for vote := range voteChannel {
		logrus.Debugf("New vote incoming: %v", vote)
		server.RconText(false, vote.PlayerId, "New vote incoming: %v", vote)
	}
}