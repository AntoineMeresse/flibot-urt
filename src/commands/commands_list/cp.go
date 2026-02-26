package commandslist

import appcontext "github.com/AntoineMeresse/flibot-urt/src/context"

func Cp(cmd *appcontext.CommandsArgs) {
	enabled := cmd.Context.Runs.ToggleCp(cmd.PlayerId)
	if enabled {
		cmd.RconText("^7Compare checkpoints: ^2On ^7(with best time in database)")
	} else {
		cmd.RconText("^7Compare checkpoints: ^1Off")
	}
}
