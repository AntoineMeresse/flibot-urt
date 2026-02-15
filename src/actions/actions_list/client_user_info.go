package actionslist

import (
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"

	log "github.com/sirupsen/logrus"

	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func ClientUserinfo(actionParams []string, c *appcontext.AppContext) {
	log.Debugf("Client User Info: %v", actionParams)
	if len(actionParams) > 1 {
		playerNumber := actionParams[0]
		infoString := strings.Join(actionParams[1:], "")
		info := splitInfos(infoString)

		if guid, ok := info["cl_guid"]; ok {
			currentPlayer := c.Players.PlayerMap[playerNumber]

			if currentPlayer == nil {
				player := c.DB.GetPlayerByGuid(guid)
				currentPlayer = &player
				c.Players.AddPlayer(playerNumber, currentPlayer)
				log.Debugf("Player %s not found. Creating it (%v)", playerNumber, player)
				c.RconText(false, playerNumber, "^4Welcome back on server. This is a ^1test server^4 so some features might be ^1broken^4.")
			}

			// Only player update
			wasUpdated := c.Players.UpdatePlayer(currentPlayer, info)
			if wasUpdated {
				log.Infof("Need to update db with new player info: %v", *currentPlayer)
			}
		} else {
			log.Warn("Could not find guid in client user info")
			// c.Players.UpdatePlayer(playerNumber, info, c.DB.GetPlayerByGuid("not_found"))
		}
	}
}

func splitInfos(infos string) map[string]string {
	res := make(map[string]string)
	infoSplit := utils.CleanEmptyElements(strings.Split(infos, "\\"))
	for i := 0; i < len(infoSplit)-1; i += 2 {
		res[infoSplit[i]] = infoSplit[i+1]
	}
	return res
}
