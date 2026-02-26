package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Aliases(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconUsage()
		return
	}

	target, err := cmd.Context.Players.GetPlayer(cmd.Params[0])
	if err != nil {
		cmd.RconText("%s", err.Error())
		return
	}

	if len(target.Aliases) == 0 {
		cmd.RconText("^7No aliases found for ^5%s^7.", target.Name)
		return
	}

	cmd.RconText("^7Aliases for ^5%s ^7(%d):", target.Name, len(target.Aliases))
	for _, chunk := range utils.ToShorterChunkArraySep(target.Aliases, ", ", false) {
		cmd.RconText("^3%s", chunk)
	}
}
