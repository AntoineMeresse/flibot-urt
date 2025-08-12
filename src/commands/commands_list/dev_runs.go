package commandslist

import (
	"fmt"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func RunsDescribe(cmd *appcontext.CommandsArgs) {
	cmd.RconText("Run describe: ")
	for k, v := range cmd.Context.Runs.PlayerRuns {
		cmd.RconText(fmt.Sprintf("----> %s %v", k, v))
	}
}

func RunsHistory(cmd *appcontext.CommandsArgs) {
	cmd.RconText("Run history: ")
	for k, v := range cmd.Context.Runs.History {
		cmd.RconText(fmt.Sprintf("----> %s %v", k, v))
	}
}
