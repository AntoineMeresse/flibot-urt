package voteslist

import (
	"fmt"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func MapVote(param string, c *appcontext.AppContext) {
	c.RconCommand("map %s", param)
}

func MapMessage(c *appcontext.AppContext, msg string, param string) (bool, string) {
	if !c.IsMapAlreadyDownloaded(param) {
		c.RconText(true, "", "^3No map found using (^6%s^3)", param)
		return false, ""
	}
	return true, fmt.Sprintf(msg, param)
}
