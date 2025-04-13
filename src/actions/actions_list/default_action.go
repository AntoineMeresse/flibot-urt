package actionslist

import (
	log "github.com/sirupsen/logrus"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func DefaultAction(actionParams []string, context *models.Context) {
	log.Debugf("DefaultAction: %v", actionParams)
}
