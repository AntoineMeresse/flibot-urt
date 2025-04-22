package actionslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func ClientBegin(actionParams []string, _ *context.Context) {
	log.Debugf("Client Begin: %v", actionParams)
}
