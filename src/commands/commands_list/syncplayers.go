package commandslist

import (
	"log/slog"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func SyncPlayers(cmd *appcontext.CommandsArgs) {
	cmd.RconText("Test sync")
	res := cmd.Context.Rcon.RconCommand("players")
	slog.Debug("Sync players", "result", res)
}
