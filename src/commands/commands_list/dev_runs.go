package commandslist

import (
	"fmt"
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func RunsDescribe(cmd *models.CommandsArgs) {
	cmd.RconText("Run describe: ")
	for k, v := range cmd.Context.Runs.PlayerRuns {
		cmd.RconText(fmt.Sprintf("----> %s %v", k, v))
	}
}

func RunsHistory(cmd *models.CommandsArgs) {
	cmd.RconText("Run history: ")
	for k, v := range cmd.Context.Runs.History {
		cmd.RconText(fmt.Sprintf("----> %s %v", k, v))
	}
}
