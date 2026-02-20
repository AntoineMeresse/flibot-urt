package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/sirupsen/logrus"
)

func ResetOptions(cmd *appcontext.CommandsArgs) {
	mapname := cmd.Context.GetCurrentMap()
	if _, err := cmd.Context.DB.DeleteMapOptions(mapname); err != nil {
		logrus.Errorf("ResetOptions DeleteMapOptions error: %v", err)
	}

	options := cmd.Context.UrtConfig.ResetOptions
	for _, opt := range options {
		cmd.RconCommand("%s", opt)
	}
	cmd.RconText("^7Reset options applied (^5%d^7)", len(options))
}
