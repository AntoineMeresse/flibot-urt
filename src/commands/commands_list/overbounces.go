package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Overbounces(cmd *models.CommandsArgs) {
	if len(cmd.Params) == 1 {
		value := cmd.Params[0]
		if value == "0" || value == "1" {
			cmd.RconCommand("g_overbounces %s", value)
			cmd.RconText("^7g_overbounces set to %s", value)
			return
		}
	}
	v := cmd.RconCommandExtractValue("g_overbounces")
	cmd.RconUsageWithText("Current value is: %s", v);
}