package actions

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func HandleAction(workerId int, action string, actionParams []string, c *appcontext.AppContext) {
	// log.Debugf("[Worker %d] ", workerId)
	log.Debugf("")
	log.Debugf("-------------------------------------------------------------------------------------------------------------")
	if val, ok := Actions[action]; ok {
		val.(func([]string, *appcontext.AppContext))(actionParams, c)
	} else {
		log.Errorf("----> Not a known action: %s\n", action)
	}
}
