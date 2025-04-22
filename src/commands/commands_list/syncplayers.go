package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/sirupsen/logrus"
)

func SyncPlayers(cmd *models.CommandsArgs) {
	cmd.RconText("Test sync")
	res := cmd.Context.Rcon.RconCommand("players")
	logrus.Debugf("Sync players: %s", res)
}
