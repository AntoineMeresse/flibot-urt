package logs

import (
	"strings"

	"github.com/AntoineMeresse/flibot-urt/src/actions"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func HandleLogsWorker(myLogChannel <-chan string, id int, server models.Server) {
	for log := range myLogChannel {
		logSplit := utils.CleanEmptyElements(strings.Split(strings.TrimSpace(log), " "))
		if len(logSplit) >= 4 {
			actions.HandleAction(id, logSplit[1], logSplit[2:], server)
		}
	}
}