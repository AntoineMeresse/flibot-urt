package actionslist

import (
	"encoding/json"
	"strconv"
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	log "github.com/sirupsen/logrus"
)

func PlayersDump(actionParams []string, c *appcontext.AppContext) {
	log.Tracef("PLayersDump action | (%d)", len(actionParams))
	if len(actionParams) >= 1 {
		dump := strings.Join(actionParams, " ")
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
