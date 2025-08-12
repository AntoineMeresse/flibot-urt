package voteslist

import (
	"fmt"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func MapVote(param string, c *appcontext.AppContext) {
	c.RconCommand("map %s", param)
}

func MapMessage(c *appcontext.AppContext, msg string, param string) (bool, string) {
	mapname, err := c.GetMapWithCriteria(param)
	if err != nil {
		c.RconText(true, "", err.Error())
		return false, ""
	} else {
		return true, fmt.Sprintf(msg, *mapname)
	}
}
