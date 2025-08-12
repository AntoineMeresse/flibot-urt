package voteslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Nextmap(param string, c *appcontext.AppContext) {
	c.RconCommand("g_nextmap %s", param)
}
