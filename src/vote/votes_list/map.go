package voteslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func MapVote(param string, context *models.Context) {
	context.RconCommand("map %s", param)
}

func MapMessage(context *models.Context, msg string, param string) (bool, string) {
	mapname, err := context.GetMapWithCriteria(param)
	if err != nil {
		context.RconText(true, "", err.Error())
		return false, ""
	} else {
		return true, fmt.Sprintf(msg, *mapname)
	}
}