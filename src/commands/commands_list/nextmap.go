package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func NextMap(cmd *models.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconText(cmd.Context.GetNextMap())
		return
	} 
	ChangeNextMap(cmd)
}