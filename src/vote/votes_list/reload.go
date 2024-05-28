package voteslist

import "github.com/AntoineMeresse/flibot-urt/src/models"

func Reload(param string, server *models.Context) {
	server.Rcon.RconCommand("reload")
}