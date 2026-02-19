package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/sirupsen/logrus"
)

func RemoveMapOptions(cmd *appcontext.CommandsArgs) {
	mapname := cmd.Context.GetCurrentMap()
	deleted, err := cmd.Context.DB.DeleteMapOptions(mapname)
	if err != nil {
		logrus.Errorf("DeleteMapOptions error: %v", err)
		return
	}
	if deleted {
		cmd.RconText("^5%s^3 options deleted.", mapname)
	} else {
		cmd.RconText("^5%s^3 has no options to delete.", mapname)
	}
}
