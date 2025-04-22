package actionslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	log "github.com/sirupsen/logrus"
)

func TimelimitHit(actionParams []string, c *context.Context) {
	log.Debugf("Timelimit hit: %v", actionParams)
	v := models.Vote{Params: []string{"extend"}}
	c.NewVote(v)
}
