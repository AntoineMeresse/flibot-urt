package actionslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func DefaultAction(actionParams []string, _ *appcontext.AppContext) {
	log.Debugf("DefaultAction: %v", actionParams)
}
