package actionslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func Tell(actionParams []string, _ *appcontext.AppContext) {
	log.Debugf("Tell: %v", actionParams)
}
