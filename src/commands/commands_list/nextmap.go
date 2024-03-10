package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func NextMap(server *models.Server, playerNumber string, params []string, isGlobal bool) {
	server.RconText(server.GetNextMap(), isGlobal, playerNumber)
}