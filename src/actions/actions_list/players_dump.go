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
				c.InitPlayerFromDump(playerNumber, p)
			}
		}
	}
}

func convertPlayersDump(line string) ([]models.DumpPlayer, error) {
	var players []models.DumpPlayer
	err := json.Unmarshal([]byte(line), &players)
	return players, err
}
