package actionslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func DefaultAction(actionParams []string, _ *context.Context) {
	log.Debugf("DefaultAction: %v", actionParams)
}
