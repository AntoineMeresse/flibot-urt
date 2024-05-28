package voteslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Cyclemap(param string, context *models.Context) {
	context.Rcon.RconCommand("cyclemap")
	// set currentmap
}

func CyclemapMessage(context *models.Context, msg string, param string) (bool, string) {
	return true, fmt.Sprintf(msg, context.GetNextMap())
}