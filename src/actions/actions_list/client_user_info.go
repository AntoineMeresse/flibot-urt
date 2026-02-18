package actionslist

import (
	"log/slog"
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"

	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func ClientUserinfo(actionParams []string, c *appcontext.AppContext) {
	slog.Debug("Client User Info", "params", actionParams)
	if len(actionParams) > 1 {
		playerNumber := actionParams[0]
		infoString := strings.Join(actionParams[1:], "")
		info := splitInfos(infoString)

		if guid, ok := info["cl_guid"]; ok {
			currentPlayer := c.Players.PlayerMap[playerNumber]

			if currentPlayer == nil {
				player, _ := c.DB.GetPlayerByGuid(guid)
				currentPlayer = &player

				c.Players.AddPlayer(playerNumber, currentPlayer)
				slog.Debug("Player not found. Creating it", "number", playerNumber, "player", player)
				c.RconText(false, playerNumber, "^4Welcome back on server. This is a ^1test server^4 so some features might be ^1broken^4.")
			}

			// Only player update
			wasUpdated := c.Players.UpdatePlayer(currentPlayer, info)
			if wasUpdated {
				slog.Info("Need to update db with new player info", "player", *currentPlayer)
			}
		} else {
			slog.Warn("Could not find guid in client user info")
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
