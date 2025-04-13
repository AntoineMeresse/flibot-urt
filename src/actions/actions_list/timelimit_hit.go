package actionslist

import (
	log "github.com/sirupsen/logrus"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func TimelimitHit(actionParams []string, context *models.Context) {
	log.Debugf("Timelimit hit: %v", actionParams)
	v := models.Vote{Params: []string{"extend"}}
	context.NewVote(v)
}
