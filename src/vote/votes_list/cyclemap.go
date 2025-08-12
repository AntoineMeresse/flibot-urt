package voteslist

import (
	"fmt"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Cyclemap(_ string, c *appcontext.AppContext) {
	c.Rcon.RconCommand("cyclemap")
	// set currentmap
}

func CyclemapMessage(c *appcontext.AppContext, msg string, _ string) (bool, string) {
	return true, fmt.Sprintf(msg, c.GetNextMap())
}
