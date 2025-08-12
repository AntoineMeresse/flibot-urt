package voteslist

import (
	"fmt"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"

	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Extend(param string, c *appcontext.AppContext) {
	c.RconCommand("extend %d", getTime(param))
}

func ExtendMessage(_ *appcontext.AppContext, msg string, param string) (bool, string) {
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
