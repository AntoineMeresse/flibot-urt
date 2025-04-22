package voteslist

import (
	"fmt"
	"github.com/AntoineMeresse/flibot-urt/src/context"

	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Extend(param string, c *context.Context) {
	c.RconCommand("extend %d", getTime(param))
}

func ExtendMessage(_ *context.Context, msg string, param string) (bool, string) {
	return true, fmt.Sprintf(msg, getTime(param))
}

func getTime(param string) int {
	extendTime := 60
	if param != "" {
		if v, err := utils.ExtractNumber(param); err == nil {
			extendTime = v
		}
	}
	return extendTime
}
