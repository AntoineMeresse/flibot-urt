package commandslist

import (
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Quit(cmd *appcontext.CommandsArgs) {
	reason := "This player isn't good enough for this map !"
	if len(cmd.Params) > 0 {
		reason = strings.Join(cmd.Params, " ")
	}
	cmd.RconCommand("kick %s \"%s\"", cmd.PlayerId, reason)
}
