package actions

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func HandleAction(workerId int, action string, actionParams []string, c *context.Context) {
	// log.Debugf("[Worker %d] ", workerId)
	if val, ok := Actions[action]; ok {
		val.(func([]string, *context.Context))(actionParams, c)
	} else {
		log.Errorf("----> Not a known action: %s\n", action)
	}
}
