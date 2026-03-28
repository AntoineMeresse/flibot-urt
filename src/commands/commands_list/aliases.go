package commandslist

import (
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Aliases(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconUsage()
		return
	}

	if strings.HasPrefix(cmd.Params[0], "@") {
		r, ok := cmd.ResolveAtId(cmd.Params[0])
		if !ok {
			return
		}
		displayAliases(cmd, r.Name, strings.Split(r.Aliases, ","))
		return
	}

	target, err := cmd.Context.Players.GetPlayer(cmd.Params[0])
	if err != nil {
		cmd.RconText("%s", err.Error())
		return
	}

	displayAliases(cmd, target.Name, target.Aliases)
}

func displayAliases(cmd *appcontext.CommandsArgs, name string, aliases []string) {
	if len(aliases) == 0 {
		cmd.RconText("^7No aliases found for ^5%s^7.", name)
		return
	}
	cmd.RconText("^7Aliases for ^5%s ^7(%d):", name, len(aliases))
	for _, chunk := range utils.ToShorterChunkArraySep(aliases, ", ", false) {
		cmd.RconText("^3%s", chunk)
	}
}
