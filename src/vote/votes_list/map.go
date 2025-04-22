package voteslist

import (
	"fmt"
	"github.com/AntoineMeresse/flibot-urt/src/context"
)

func MapVote(param string, c *context.Context) {
	c.RconCommand("map %s", param)
}

func MapMessage(c *context.Context, msg string, param string) (bool, string) {
	mapname, err := c.GetMapWithCriteria(param)
	if err != nil {
		c.RconText(true, "", err.Error())
		return false, ""
	} else {
		return true, fmt.Sprintf(msg, *mapname)
	}
}
