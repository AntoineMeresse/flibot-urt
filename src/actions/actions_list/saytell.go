package actionslist

import (
	"log/slog"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func SayTell(actionParams []string, _ *appcontext.AppContext) {
	slog.Debug("SayTell", "params", actionParams)
}
