package voteslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
)

func Nextmap(param string, c *context.Context) {
	c.RconCommand("g_nextmap %s", param)
}
