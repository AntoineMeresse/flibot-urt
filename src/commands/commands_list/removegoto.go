package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils/msg"
	"github.com/sirupsen/logrus"
)

func RemoveGoto(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) > 0 {
		jumpName := cmd.Params[0]
		mapname := cmd.Context.GetCurrentMap()
		deleted, err := cmd.Context.DB.DeleteGoto(mapname, jumpName)
		if err != nil {
			logrus.Errorf("DeleteGoto error: %v", err)
			return
		}
		if deleted {
			cmd.RconText(msg.GOTO_REMOVE, jumpName)
		} else {
			cmd.RconText(msg.GOTO_DONT_EXIST, jumpName)
		}
	}
}
