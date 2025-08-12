package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/sirupsen/logrus"
)

func SyncPlayers(cmd *appcontext.CommandsArgs) {
	cmd.RconText("Test sync")
	res := cmd.Context.Rcon.RconCommand("players")
	logrus.Debugf("Sync players: %s", res)
}
