package logs

import (
	"log/slog"
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"

	"github.com/AntoineMeresse/flibot-urt/src/actions"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func HandleLogsWorker(myLogChannel <-chan string, id int, c *appcontext.AppContext) {
	slog.Debug("Init worker", "id", id)
	for line := range myLogChannel {
		slog.Debug("Worker read", "line", line)
		logSplit := utils.CleanEmptyElements(strings.Split(strings.TrimSpace(line), " "))
		slog.Debug("Log:", "parts", logSplit)
		if len(logSplit) >= 3 {
			slog.Debug("Log Ok:", "parts", logSplit)
			actions.HandleAction(id, logSplit[1], logSplit[2:], c)
		}
	}
}
