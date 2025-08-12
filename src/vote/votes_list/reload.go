package voteslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Reload(_ string, c *appcontext.AppContext) {
	c.Rcon.RconCommand("reload")
}
