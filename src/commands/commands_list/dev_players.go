package commandslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
	log "github.com/sirupsen/logrus"
)

func PlayersList(cmd *models.CommandsArgs) {
	cmd.RconText("Players: ")
	cmd.Context.Players.Mutex.RLock()
	for key, value := range cmd.Context.Players.PlayerMap {
		text := fmt.Sprintf("%s: %v", key, value)
		log.Debug(text)
		cmd.RconText(text)
	}
	cmd.Context.Players.Mutex.RUnlock()
}

func PlayersGet(cmd *models.CommandsArgs) {
	if len(cmd.Params) > 0 {
		searchCriteria := cmd.Params[0]
		cmd.RconText("Player with criteria (%s): ", searchCriteria)
		player, err := cmd.Context.Players.GetPlayer(searchCriteria)
		if err == nil {
			cmd.RconText("Player found: (%v)", *player)
		} else {
			cmd.RconText(err.Error())
		}
	} else {
		cmd.RconText("Add criteria")
	}
}
