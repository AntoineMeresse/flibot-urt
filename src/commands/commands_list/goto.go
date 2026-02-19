package commandslist

import (
	gotoshared "github.com/AntoineMeresse/flibot-urt/src/commands/shared/goto"
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/AntoineMeresse/flibot-urt/src/utils/msg"
	"github.com/sirupsen/logrus"
)

func Goto(cmd *appcontext.CommandsArgs) {
	mapname := cmd.Context.GetCurrentMap()
	if len(cmd.Params) == 0 {
		names, err := cmd.Context.DB.GetGotoNames(mapname)
		if err != nil {
			logrus.Errorf("GetGotoNames error: %v", err)
			return
		}
		cmd.RconList(gotoshared.BuildDisplayLocation(mapname, utils.NaturalSort(names)))
	} else {
		jumpName := cmd.Params[0]
		g, ok := cmd.Context.DB.GetGoto(mapname, jumpName)
		if !ok {
			cmd.RconText(msg.GOTO_NO_LOCATION, jumpName)
			return
		}
		cmd.RconCommand("forceteam %s free", cmd.PlayerId)
		cmd.RconCommand("tpPos %s %g %g %g %g %g", cmd.PlayerId, g.PosX, g.PosY, g.PosZ, g.AngleV, g.AngleH)
	}
}
