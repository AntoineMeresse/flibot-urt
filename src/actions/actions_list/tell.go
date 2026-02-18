package actionslist

import (
	"log/slog"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Tell(actionParams []string, _ *appcontext.AppContext) {
	slog.Debug("Tell", "params", actionParams)
}
