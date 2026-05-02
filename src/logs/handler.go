package logs

import (
	"regexp"
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"

	"github.com/AntoineMeresse/flibot-urt/src/actions"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/sirupsen/logrus"
)

var timestampRegex = regexp.MustCompile(`^\s*\d+:\d+\s*`)

func HandleLogsWorker(myLogChannel <-chan string, id int, c *appcontext.AppContext) {
	logrus.Tracef("Init worker: %d", id)
	for log := range myLogChannel {
		logrus.Tracef("Worker %d read raw: %q", id, log)

		cleanLog := timestampRegex.ReplaceAllString(log, "")  // Strip timestamp (e.g., "1308:56ClientConnect:" -> "ClientConnect:")
		logSplit := utils.CleanEmptyElements(strings.Split(strings.TrimSpace(cleanLog), " "))
		logrus.Tracef("Worker %d split: %v", id, logSplit)

		if len(logSplit) >= 1 {
			action := logSplit[0]
			params := logSplit[1:]
			actions.HandleAction(id, action, params, c)
		} else {
			logrus.Errorf("Worker %d: empty line after cleaning, skipping", id)
		}
	}
}
