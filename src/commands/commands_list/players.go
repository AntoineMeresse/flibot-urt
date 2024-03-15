package commandslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
	log "github.com/sirupsen/logrus"
)

func PlayersList(cmd *models.CommandsArgs) {
	cmd.RconText("Players: ")
	cmd.Server.Players.Mutex.RLock()
	for key, value := range(cmd.Server.Players.List) {
		text := fmt.Sprintf("%s: %v" , key,  value)
		log.Debug(text)
		cmd.RconText(text)
	}
	cmd.Server.Players.Mutex.RUnlock()
}

func PlayersGet(cmd *models.CommandsArgs) {
	if len(cmd.Params) > 0 {
		searchCriteria := cmd.Params[0]
		cmd.RconText(fmt.Sprintf("Player with criteria (%s): ", searchCriteria))
		player, err := cmd.Server.Players.GetPlayer(searchCriteria)
		if err == nil {
			text := fmt.Sprintf("Player found: (%v)" , *player)
			cmd.RconText(text)
		} else {
			cmd.RconText(err.Error())
		}
	} else {
		cmd.RconText("Add criteria")
	}
}