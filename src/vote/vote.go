package vote

import (
	"log/slog"
	"strings"
	"sync"
	"time"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"

	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

const (
	SecondsPerVote = 20
)

type VoteSystem struct {
	mu      sync.RWMutex
	CanVote bool
	Cancel  bool
	VoteYes map[string]int // use of map cause golang has no set
	VoteNo  map[string]int // use of map cause golang has no set
}

func InitVoteSystem(voteChannel <-chan models.Vote, c *appcontext.AppContext) {
	voteSystem := VoteSystem{CanVote: true, Cancel: false, VoteYes: make(map[string]int), VoteNo: make(map[string]int)}
	slog.Debug("VoteSystem initiated", "system", &voteSystem)

	for vote := range voteChannel {
		slog.Debug("New vote incoming", "vote", vote)
		go voteLogic(c, &voteSystem, vote)
	}
}

func (voteSystem *VoteSystem) reset() {
	voteSystem.mu.Lock()
	defer voteSystem.mu.Unlock()
	voteSystem.CanVote = true
	voteSystem.Cancel = false
	clear(voteSystem.VoteYes)
	clear(voteSystem.VoteNo)
}

func voteLogic(c *appcontext.AppContext, voteSystem *VoteSystem, vote models.Vote) {
	if vote.Params == nil {
		slog.Error("Vote params can't be null", "vote", vote)
		return
	}

	if handleVote(voteSystem, vote) {
		return
	}

	createVote(c, voteSystem, vote)
}

func isOnlyVote(vote models.Vote) (isVote bool, value string) {
	return len(vote.Params) == 1 && utils.IsVoteCommand(vote.Params[0]), vote.Params[0]
}

func createVote(c *appcontext.AppContext, voteSystem *VoteSystem, vote models.Vote) {
	if !voteSystem.tryClaimVote() {
		c.RconText(false, vote.PlayerId, "Can't ^1start^3 a new vote !")
		return
	}

	continueVote, endFunction, msg := getVoteInfos(c, vote)
	if !continueVote {
		voteSystem.reset()
		return
	}

	c.RconText(false, vote.PlayerId, "New vote incoming: %v", vote)
	iteration := 0
	secondsToEnd := SecondsPerVote * 2 // To avoid to deal with float
	cpt := 0
	for iteration <= secondsToEnd && !voteSystem.isCanceled() {
		voteKeysMessage(&cpt, c)
		yes, no := voteSystem.voteCounts()
		c.RconBigText("%s | ^2Yes^7: %2d / ^1No^7 : %2d (%02d s)", msg, yes, no, (secondsToEnd-iteration)/2)
		iteration += 1
		time.Sleep(500 * time.Millisecond)
		if hasMajority(c, voteSystem) {
			break
		}
	}
	endVote(c, voteSystem, vote, endFunction)
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
		return true
	}
	return false
}

func (voteSystem *VoteSystem) addYesVote(playerId string) {
	voteSystem.mu.Lock()
	defer voteSystem.mu.Unlock()
	delete(voteSystem.VoteNo, playerId)
	voteSystem.VoteYes[playerId] = 0
}

func (voteSystem *VoteSystem) addNoVote(playerId string) {
	voteSystem.mu.Lock()
	defer voteSystem.mu.Unlock()
	delete(voteSystem.VoteYes, playerId)
	voteSystem.VoteNo[playerId] = 0
}

func (voteSystem *VoteSystem) addVetoVote() {
	voteSystem.mu.Lock()
	defer voteSystem.mu.Unlock()
	voteSystem.Cancel = true
}

func voteKeysMessage(cpt *int, c *appcontext.AppContext) {
	if *cpt == 10 {
		*cpt = 0
	}
	if *cpt == 0 {
		c.RconPrint("^7Use [^2'+'^7] or [^1'-'^7] to vote.")
	}
	*cpt += 2
}

func hasMajority(c *appcontext.AppContext, voteSystem *VoteSystem) bool {
	majority := (len(c.Players.PlayerMap) / 2) + 1
	yes, no := voteSystem.voteCounts()
	return yes >= majority || no >= majority
}

func endVote(c *appcontext.AppContext, voteSystem *VoteSystem, vote models.Vote, endFunction interface{}) {
	canceled, yesWins := voteSystem.outcome()
	if !canceled {
		if yesWins {
			c.RconBigText("^2Vote Passed")
			execVote(c, vote, endFunction)
		} else {
			c.RconBigText("^1Vote Failed")
		}
	} else {
		c.RconBigText("^1Vote Canceled")
	}
	voteSystem.reset()
}

func getVoteInfos(c *appcontext.AppContext, vote models.Vote) (bool, interface{}, string) {
	infos, exists := Votes[vote.Params[0]]
	param := strings.Join(vote.Params[1:], " ")
	if exists {
		continueVote, msg := infos.msgFn.(func(*appcontext.AppContext, string, string) (bool, string))(c, infos.messageFormat, param)
		return continueVote, infos.function, msg
	}
	c.RconText(false, vote.PlayerId, "Vote [%s] does not exist", vote.Params[0])
	return false, nil, ""
}

func execVote(c *appcontext.AppContext, vote models.Vote, endFunction interface{}) {
	time.Sleep(1 * time.Second)
	param := strings.Join(vote.Params[1:], " ")
	endFunction.(func(string, *appcontext.AppContext))(param, c)
}

func (voteSystem *VoteSystem) tryClaimVote() bool {
	voteSystem.mu.Lock()
	defer voteSystem.mu.Unlock()
	if !voteSystem.CanVote {
		return false
	}
	voteSystem.CanVote = false
	return true
}

func (voteSystem *VoteSystem) isCanceled() bool {
	voteSystem.mu.RLock()
	defer voteSystem.mu.RUnlock()
	return voteSystem.Cancel
}

func (voteSystem *VoteSystem) voteCounts() (yes, no int) {
	voteSystem.mu.RLock()
	defer voteSystem.mu.RUnlock()
	return len(voteSystem.VoteYes), len(voteSystem.VoteNo)
}

func (voteSystem *VoteSystem) outcome() (canceled bool, yesWins bool) {
	voteSystem.mu.RLock()
	defer voteSystem.mu.RUnlock()
	return voteSystem.Cancel, len(voteSystem.VoteYes) > len(voteSystem.VoteNo)
}
