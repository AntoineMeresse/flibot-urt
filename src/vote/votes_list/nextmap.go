package voteslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Nextmap(param string, context *models.Context) {
	context.RconCommand("g_nextmap %s", param)
}
