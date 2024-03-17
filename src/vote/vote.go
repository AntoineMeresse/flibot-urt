package vote

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/sirupsen/logrus"
)

type VoteSystem struct {
	CanVote bool
	Cancel  bool
	VoteYes map[string]int // use of map cause golang has no set
	VoteNo  map[string]int // use of map cause golang has no set
}

func InitVoteSystem(voteChannel <-chan models.Vote, server *models.Server) {
	voteSystem := &VoteSystem{CanVote: true, Cancel: false, VoteYes: make(map[string]int), VoteNo: make(map[string]int)}
	logrus.Debugf("VoteSystem initiated: %v", voteSystem)

	for vote := range voteChannel {
		logrus.Debugf("New vote incoming: %v", vote)
		go voteLogic(server, voteSystem, vote)
	}
}

func voteLogic(server *models.Server, voteSystem *VoteSystem, vote models.Vote) {
	if vote.Params == nil {
		logrus.Errorf("Vote params can't be null %v", vote)
		return
	}

	if handleVote(server, voteSystem, vote) {
		return
	} 
	
	createVote(server, voteSystem, vote)
}

func isOnlyVote(vote models.Vote) (isVote bool, value string) {
	return len(vote.Params) == 1 && utils.IsVoteCommand(vote.Params[0]), vote.Params[0]
}

func createVote(server *models.Server, voteSystem *VoteSystem, vote models.Vote) {
	if (voteSystem.CanVote) {
		server.RconText(false, vote.PlayerId, "New vote incoming: %v", vote)
		voteSystem.CanVote = false
	} else {
		server.RconText(false, vote.PlayerId, "Can't start a new vote !")
	}
}

func handleVote(server *models.Server, voteSystem *VoteSystem, vote models.Vote) (isVote bool) {
	if isVote, value := isOnlyVote(vote); isVote {
		server.RconText(false, vote.PlayerId, "Just a vote !")
		if (value == "+") {
			voteSystem.addYesVote(vote.PlayerId)
		} else {
			voteSystem.addNoVote(vote.PlayerId)
		}
		logrus.Debugf("Vote system values %v", *voteSystem)
		return true
	}
	return false
}


func (v *VoteSystem) addYesVote(playerId string) {
	v.VoteYes[playerId] = 0
}

func (v *VoteSystem) addNoVote(playerId string) {
	v.VoteNo[playerId] = 0
}