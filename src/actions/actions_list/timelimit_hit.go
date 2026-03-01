package actionslist

import (
	"math/rand"
	"time"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	log "github.com/sirupsen/logrus"
)

var runningMessages = []string{
	"I'm currently running, please wait !",
	"Don't you dare end this map, I'm on a run !",
	"EXTEND ! I can feel it this time !",
	"I was BORN to finish this run !",
	"My fingers are sweating, extend pls",
	"One more try I swear",
	"I'm so close, don't kill my dreams",
	"If you don't extend I will cry",
}

func TimelimitHit(actionParams []string, c *appcontext.AppContext) {
	log.Debugf("Timelimit hit: %v", actionParams)
	if c.VoteActive.Load() {
		c.RconText(true, "", "^7A vote is still active, extend skipped.")
		return
	}
	v := models.Vote{Params: []string{"extend"}}
	c.NewVote(v)
	go announceRunningPlayers(c)
}

func announceRunningPlayers(c *appcontext.AppContext) {
	time.Sleep(500 * time.Millisecond)
	numbers := c.Runs.RunningPlayerNumbers()
	for _, number := range numbers {
		msg := runningMessages[rand.Intn(len(runningMessages))]
		c.RconCommand("spoof %s say + %s", number, msg)
	}
}
