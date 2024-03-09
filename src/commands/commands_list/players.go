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
	server.RconText("Players: ", isGlobal, playerNumber)
	if len(params) > 0 {
		searchCriteria := params[0]
		found, player := server.Players.GetPlayer(searchCriteria)
		if found {
			text := fmt.Sprintf("Player found with these datas: (%v)" , player)
			server.RconText(text, isGlobal, playerNumber)
		} else {
			server.RconText("No player found", isGlobal, playerNumber)
		}
	} else {
		server.RconText("Add criteria", isGlobal, playerNumber)
	}
}