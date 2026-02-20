package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Overbounces(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 1 {
		value := cmd.Params[0]
		if value == "0" || value == "1" {
			cmd.RconCommand("g_overbounces %s", value)
			cmd.RconText("^7g_overbounces set to %s", value)
			return
		}
	}
	rcon := NewRconClient(cmd)
	defer rcon.CloseConnection()
	v := rcon.RconCommandExtractValue("g_overbounces")
	cmd.RconUsageWithText("Current value is: %s", v)
}
