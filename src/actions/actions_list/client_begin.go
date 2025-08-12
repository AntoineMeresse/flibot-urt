package actionslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func ClientBegin(actionParams []string, _ *appcontext.AppContext) {
	log.Debugf("Client Begin: %v", actionParams)
}
