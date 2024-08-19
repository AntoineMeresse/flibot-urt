package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func SetNextMap(cmd *models.CommandsArgs) {
	ChangeNextMap(cmd)
}

func ChangeNextMap(cmd *models.CommandsArgs) {
	if len(cmd.Params) >= 1 {
		mapName, err := cmd.Context.GetMapWithCriteria(cmd.Params[0])
		if err != nil {
			cmd.RconText("Can not change next map because: ^7%s", err.Error())
		} else {
			cmd.RconText("Changing nextmap to: %s", *mapName)
			cmd.Context.SetNextMap(*mapName)
		}
	}
}