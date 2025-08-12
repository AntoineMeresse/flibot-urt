package actionslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	log "github.com/sirupsen/logrus"
)

func TimelimitHit(actionParams []string, c *appcontext.AppContext) {
	log.Debugf("Timelimit hit: %v", actionParams)
	v := models.Vote{Params: []string{"extend"}}
	c.NewVote(v)
}
