package actionslist

import (
	"log/slog"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func DefaultAction(actionParams []string, _ *appcontext.AppContext) {
	slog.Debug("DefaultAction", "params", actionParams)
}
