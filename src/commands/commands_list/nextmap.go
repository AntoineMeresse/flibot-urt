package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func NextMap(server models.Server, playerNumber string, params []string, isGlobal bool) {
	server.RconTextInfo(server.Nextmap, isGlobal, playerNumber)
}