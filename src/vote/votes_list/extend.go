package voteslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Extend(param string, server *models.Context) {
	server.RconCommand("extend %d", getTime(param))
}

func ExtendMessage(server *models.Context, msg string, param string) (bool, string) {
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