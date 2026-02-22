package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Nuke(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) > 0 {
		p, err := cmd.Context.Players.GetPlayer(cmd.Params[0])
		if err != nil {
			cmd.RconText("%s", err.Error())
			return
		}

		cmd.RconCommand("nuke %s", p.Number)
		cmd.RconText("^5%s^7 was nuked !", p.Name)
	} else {
		cmd.RconUsage()
	}
}
