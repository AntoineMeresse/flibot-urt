package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Extend(cmd *models.CommandsArgs) {
	extendTime := -1;
	
	if len(cmd.Params) == 0 {
		extendTime = 60
	} else if len(cmd.Params) > 0 {
		t, err := utils.ExtractNumber(cmd.Params[0])

		if err == nil {
			if t > 0 && t < 1000 {
				extendTime = t;
			}
		} 
	}
	
	if extendTime > 0 {
		cmd.RconCommand("extend %d", extendTime)
	} else {
		cmd.RconUsage(cmd.Usage)
	}
}