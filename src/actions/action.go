package actions

import (
	"log/slog"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func HandleAction(workerId int, action string, actionParams []string, c *appcontext.AppContext) {
	slog.Debug("-------------------------------------------------------------------------------------------------------------")
	if val, ok := Actions[action]; ok {
		val(actionParams, c)
	} else {
		slog.Error("----> Not a known action:", "action", action)
	}
}
