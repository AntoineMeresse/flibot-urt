package actionslist

import (
	"log/slog"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func ClientConnect(actionParams []string, _ *appcontext.AppContext) {
	slog.Debug("Client Connect", "params", actionParams)
	//c.Players.AddPlayer(actionParams[0], &models.Player{})
}
