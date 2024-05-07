package voteslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Nextmap(param string, server *models.Server) {
	server.RconCommand("g_nextmap %s", param)
}
