package actionslist

import (
	log "github.com/sirupsen/logrus"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func TimelimitHit(action_params []string, context *models.Context) {
	log.Debugf("Timelimit hit: %v", action_params)
	v := models.Vote{Params: []string{"extend"}}
	context.NewVote(v)
}