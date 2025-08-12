package logs

import (
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"

	"github.com/AntoineMeresse/flibot-urt/src/actions"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/sirupsen/logrus"
)

func HandleLogsWorker(myLogChannel <-chan string, id int, c *appcontext.AppContext) {
	logrus.Tracef("Init worker: %d", id)
	for log := range myLogChannel {
		logrus.Tracef("Worker read: %s", log)
		logSplit := utils.CleanEmptyElements(strings.Split(strings.TrimSpace(log), " "))
		logrus.Tracef("Log: %s", logSplit)
		if len(logSplit) >= 3 {
			logrus.Tracef("Log Ok: %v", logSplit)
			actions.HandleAction(id, logSplit[1], logSplit[2:], c)
		}
	}
}
