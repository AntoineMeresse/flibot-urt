package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
)

func SetNextMap(cmd *context.CommandsArgs) {
	ChangeNextMap(cmd)
}

func ChangeNextMap(cmd *context.CommandsArgs) {
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
