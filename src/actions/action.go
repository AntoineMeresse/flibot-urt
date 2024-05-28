package actions

import (
	log "github.com/sirupsen/logrus"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func HandleAction(workerId int, action string, action_params []string, context *models.Context) {
	// log.Debugf("[Worker %d] ", workerId)
	if val, ok := Actions[action]; ok {
		val.(func([]string, *models.Context))(action_params, context)
	} else {
		log.Errorf("----> Not a known action: %s\n", action)
	}
}