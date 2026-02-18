package actionslist

import (
	"log/slog"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func ClientBegin(actionParams []string, _ *appcontext.AppContext) {
	slog.Debug("Client Begin", "params", actionParams)
}
