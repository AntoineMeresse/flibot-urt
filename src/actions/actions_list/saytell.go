package actionslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func SayTell(actionParams []string, _ *context.Context) {
	log.Debugf("SayTell: %v", actionParams)
}
