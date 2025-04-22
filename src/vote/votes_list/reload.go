package voteslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
)

func Reload(_ string, c *context.Context) {
	c.Rcon.RconCommand("reload")
}
