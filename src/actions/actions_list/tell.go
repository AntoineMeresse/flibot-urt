package actionslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func Tell(actionParams []string, _ *context.Context) {
	log.Debugf("Tell: %v", actionParams)
}
