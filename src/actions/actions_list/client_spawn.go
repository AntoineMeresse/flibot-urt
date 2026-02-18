package actionslist

import (
	"log/slog"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func ClientSpawn(actionParams []string, _ *appcontext.AppContext) {
	// When the player join in game
	slog.Debug("ClientSpawn", "params", actionParams)
}
