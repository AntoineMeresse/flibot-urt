package voteslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Cyclemap(param string, server *models.Server) {
	server.Rcon.RconCommand("cyclemap")
	// set currentmap
}

func CyclemapMessage(server *models.Server, msg string, param string) (bool, string) {
	return true, fmt.Sprintf(msg, server.GetNextMap())
}