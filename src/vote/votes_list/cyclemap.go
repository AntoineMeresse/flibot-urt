package voteslist

import (
	"fmt"
	"github.com/AntoineMeresse/flibot-urt/src/context"
)

func Cyclemap(_ string, c *context.Context) {
	c.Rcon.RconCommand("cyclemap")
	// set currentmap
}

func CyclemapMessage(c *context.Context, msg string, _ string) (bool, string) {
	return true, fmt.Sprintf(msg, c.GetNextMap())
}
