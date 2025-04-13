package vote

import (
	"strings"
	"time"

	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/sirupsen/logrus"
)

const (
	SecondsPerVote = 20
)

type VoteSystem struct {
	// Mutex needed ?
	CanVote bool
	Cancel  bool
	VoteYes map[string]int // use of map cause golang has no set
	VoteNo  map[string]int // use of map cause golang has no set
}

func InitVoteSystem(voteChannel <-chan models.Vote, context *models.Context) {
	voteSystem := VoteSystem{CanVote: true, Cancel: false, VoteYes: make(map[string]int), VoteNo: make(map[string]int)}
	logrus.Debugf("VoteSystem initiated: %v", voteSystem)

	for vote := range voteChannel {
		logrus.Debugf("New vote incoming: %v", vote)
		go voteLogic(context, &voteSystem, vote)
	}
}

func (voteSystem *VoteSystem) reset() {
	voteSystem.CanVote = true
	voteSystem.Cancel = false
	clear(voteSystem.VoteYes)
	clear(voteSystem.VoteNo)
}

func voteLogic(context *models.Context, voteSystem *VoteSystem, vote models.Vote) {
	if vote.Params == nil {
		logrus.Errorf("Vote params can't be null %v", vote)
		return
	}

	if handleVote(voteSystem, vote) {
		return
	}

	createVote(context, voteSystem, vote)
}

func isOnlyVote(vote models.Vote) (isVote bool, value string) {
	return len(vote.Params) == 1 && utils.IsVoteCommand(vote.Params[0]), vote.Params[0]
}

func createVote(context *models.Context, voteSystem *VoteSystem, vote models.Vote) {
	if voteSystem.CanVote {
		if continueVote, endFunction, msg := getVoteInfos(context, vote); continueVote {
			voteSystem.CanVote = false
			context.RconText(false, vote.PlayerId, "New vote incoming: %v", vote)
			iteration := 0
			secondsToEnd := SecondsPerVote * 2 // To avoid to deal with float
			cpt := 0
			for iteration <= secondsToEnd && !voteSystem.Cancel {
				voteKeysMessage(&cpt, context)
				context.RconBigText("%s | ^2Yes^7: %2d / ^1No^7 : %2d (%02d s)", msg, len(voteSystem.VoteYes), len(voteSystem.VoteNo), (secondsToEnd-iteration)/2)
				iteration += 1
				time.Sleep(500 * time.Millisecond)
				if hasMajority(context, voteSystem) {
					break
				}
			}
			endVote(context, voteSystem, vote, endFunction)
			voteSystem.CanVote = true
		}
	} else {
		context.RconText(false, vote.PlayerId, "Can't ^1start^3 a new vote !")
	}
}

func handleVote(voteSystem *VoteSystem, vote models.Vote) (isVote bool) {
	if isVote, value := isOnlyVote(vote); isVote {
		if value == "v" {
			voteSystem.addVetoVote()
		} else if value == "+" {
			voteSystem.addYesVote(vote.PlayerId)
		} else {
			voteSystem.addNoVote(vote.PlayerId)
		}
		logrus.Debugf("Vote system values %v", *voteSystem)
		return true
	}
	return false
}

func (voteSystem *VoteSystem) addYesVote(playerId string) {
	delete(voteSystem.VoteNo, playerId)
	voteSystem.VoteYes[playerId] = 0
}

func (voteSystem *VoteSystem) addNoVote(playerId string) {
	delete(voteSystem.VoteYes, playerId)
	voteSystem.VoteNo[playerId] = 0
}

func (voteSystem *VoteSystem) addVetoVote() {
	voteSystem.Cancel = true
}

func voteKeysMessage(cpt *int, context *models.Context) {
	if *cpt == 10 {
		*cpt = 0
	}
	if *cpt == 0 {
		context.RconPrint("^7Use [^2'+'^7] or [^1'-'^7] to vote.")
	}
	*cpt += 2
}

func hasMajority(context *models.Context, voteSystem *VoteSystem) bool {
	majority := (len(context.Players.List) / 2) + 1
	return len(voteSystem.VoteYes) >= majority || len(voteSystem.VoteNo) >= majority
}

func endVote(context *models.Context, voteSystem *VoteSystem, vote models.Vote, endFunction interface{}) {
	if !voteSystem.Cancel {
		if len(voteSystem.VoteYes) > len(voteSystem.VoteNo) {
			context.RconBigText("^2Vote Passed")
			execVote(context, vote, endFunction)
		} else {
			context.RconBigText("^1Vote Failed")
		}
	} else {
		context.RconBigText("^1Vote Canceled")
	}
	voteSystem.reset()
}

func getVoteInfos(context *models.Context, vote models.Vote) (bool, interface{}, string) {
	infos, exists := Votes[vote.Params[0]]
	param := strings.Join(vote.Params[1:], " ")
	if exists {
		continueVote, msg := infos.msgFn.(func(*models.Context, string, string) (bool, string))(context, infos.messageFormat, param)
		return continueVote, infos.function, msg
	} else {
		context.RconText(false, vote.PlayerId, "Vote [%s] does not exist", vote.Params[0])
	}
	return false, nil, ""
}

func execVote(context *models.Context, vote models.Vote, endFunction interface{}) {
	time.Sleep(1 * time.Second)
	param := strings.Join(vote.Params[1:], " ")
	endFunction.(func(string, *models.Context))(param, context)
}
