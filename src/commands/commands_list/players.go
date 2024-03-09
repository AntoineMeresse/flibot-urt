package commandslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
	log "github.com/sirupsen/logrus"
)

func PlayersList(server *models.Server, playerNumber string, params []string, isGlobal bool) {
	server.RconText("Players: ", isGlobal, playerNumber)
	server.Players.Mutex.RLock()
	for key, value := range(server.Players.List) {
		text := fmt.Sprintf("%s: %v" , key,  value)
		log.Debug(text)
		server.RconText(text, isGlobal, playerNumber)
	}
	server.Players.Mutex.RUnlock()
}

func PlayersGet(server *models.Server, playerNumber string, params []string, isGlobal bool) {
	if len(params) > 0 {
		searchCriteria := params[0]
		server.RconText(fmt.Sprintf("Player with criteria (%s): ", searchCriteria), isGlobal, playerNumber)
		player, err := server.Players.GetPlayer(searchCriteria)
		if err == nil {
			text := fmt.Sprintf("Player found: (%v)" , *player)
			server.RconText(text, isGlobal, playerNumber)
		} else {
			server.RconText(err.Error(), isGlobal, playerNumber)
		}
	} else {
		server.RconText("Add criteria", isGlobal, playerNumber)
	}
}