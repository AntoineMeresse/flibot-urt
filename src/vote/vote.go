package vote

import (
	"strings"
	"time"

	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/sirupsen/logrus"
)

const (
	SECONDS_PER_VOTE = 20
)

type VoteSystem struct {
	// Mutex needed ?
	CanVote bool
	Cancel  bool
	VoteYes map[string]int // use of map cause golang has no set
	VoteNo  map[string]int // use of map cause golang has no set
}

func InitVoteSystem(voteChannel <-chan models.Vote, server *models.Server) {
	voteSystem := VoteSystem{CanVote: true, Cancel: false, VoteYes: make(map[string]int), VoteNo: make(map[string]int)}
	logrus.Debugf("VoteSystem initiated: %v", voteSystem)

	for vote := range voteChannel {
		logrus.Debugf("New vote incoming: %v", vote)
		go voteLogic(server, &voteSystem, vote)
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
		voteSystem.CanVote = false
		server.RconText(false, vote.PlayerId, "New vote incoming: %v", vote)
		iteration := 0
		secondsToEnd := SECONDS_PER_VOTE*2 // To avoid to deal with float
		cpt := 0
		msg := strings.Join(vote.Params, " ")
		for (iteration <= secondsToEnd && !voteSystem.Cancel) {
			voteKeysMessage(&cpt, server)
			server.RconBigText(true, "0", "%s | ^2Yes^7: %2d / ^1No^7 : %2d (%02d s)", msg, len(voteSystem.VoteYes), len(voteSystem.VoteNo), (secondsToEnd - iteration) / 2)
			iteration += 1
			time.Sleep(500 * time.Millisecond)
			if hasMajority() {
				break
			}
		}
		voteSystem.CanVote = true
	} else {
		server.RconText(false, vote.PlayerId, "Can't ^1start^3 a new vote !")
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
	delete(v.VoteNo, playerId)
	v.VoteYes[playerId] = 0
}

func (v *VoteSystem) addNoVote(playerId string) {
	delete(v.VoteYes, playerId)
	v.VoteNo[playerId] = 0
}

func voteKeysMessage(cpt *int, server *models.Server,) {
	if (*cpt == 10) {
		*cpt = 0
	}
	if (*cpt == 0) {
		server.RconPrint("^7Use [^2'+'^7] or [^1'-'^7] to vote.")
	}
	*cpt += 2
}

func hasMajority() bool {
	// TODO: check if majority 
	return false
}