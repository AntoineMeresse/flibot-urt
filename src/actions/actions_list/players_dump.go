package actionslist

import (
	"encoding/json"
	"strconv"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	log "github.com/sirupsen/logrus"
)

func PlayersDump(actionParams []string, c *appcontext.AppContext) {
	if len(actionParams) == 1 {
		dump := actionParams[0]
		log.Debugf("PLayersDump: %v", dump)

		players, err := convertPlayersDump(dump)
		if err != nil {
			log.Error(err)
			return
		}

		for i, p := range players {
			log.Infof("%2d) ---> %v", i, p)

			playerNumber := strconv.Itoa(p.PlayerNumber)
			currentPlayer := c.Players.PlayerMap[playerNumber]

			if currentPlayer == nil {
				player, found := c.DB.GetPlayerByGuid(p.GUID)

				if !found {
					player = models.Player{
						Guid: p.GUID,
						Name: p.Name,
					}
				}

				currentPlayer = &player
				c.Players.AddPlayer(playerNumber, currentPlayer)
				log.Debugf("Player %s not found. Creating it (%v)", playerNumber, player)
				c.RconText(false, playerNumber, "^4Welcome back on server. This is a ^1test server^4 so some features might be ^1broken^4.")
			}
		}

	}

}

type DumpPlayer struct {
	PlayerNumber int    `json:"playernumber"`
	Name         string `json:"name"`
	GUID         string `json:"guid"`
}

func convertPlayersDump(line string) ([]DumpPlayer, error) {
	var players []DumpPlayer
	err := json.Unmarshal([]byte(line), &players)
	return players, err
}
