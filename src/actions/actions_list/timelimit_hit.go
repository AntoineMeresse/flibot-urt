package actionslist

import (
	"time"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	log "github.com/sirupsen/logrus"
)

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
		c.RconCommand("spoof %s say + [auto vote]", number)
	}
}
