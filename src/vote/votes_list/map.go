package voteslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func MapVote(param string, server *models.Server) {
	server.RconCommand("map %s", param)
}

func MapMessage(server *models.Server, msg string, param string) (bool, string) {
	mapname, err := server.GetMapWithCriteria(param)
	if err != nil {
		return false, ""
	} else {
		return true, fmt.Sprintf(msg, *mapname)
	}
}