package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Register(cmd *models.CommandsArgs) {
	cmd.RconText("Register")
	player, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	if err == nil {
		err = cmd.Context.DB.SaveNewPlayer(player.Name, player.Guid, "No ip")
		if err == nil {
			cmd.RconText("Register in db")
			return
		} 
	} 
	cmd.RconText(err.Error())
}