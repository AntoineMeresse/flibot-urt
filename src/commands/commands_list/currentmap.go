package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func CurrentMap(server *models.Server, playerNumber string, params []string, isGlobal bool) {
	server.RconText(server.Mapname, isGlobal, playerNumber)
}