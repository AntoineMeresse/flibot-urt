package voteslist

import "github.com/AntoineMeresse/flibot-urt/src/models"

func Reload(param string, context *models.Context) {
	context.Rcon.RconCommand("reload")
}