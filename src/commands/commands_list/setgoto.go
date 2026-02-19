package commandslist

import (
	gotoshared "github.com/AntoineMeresse/flibot-urt/src/commands/shared/goto"
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/quake3_rcon"
	"github.com/sirupsen/logrus"
)

func SetGoto(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) > 0 {
		res := cmd.RconCommand("playerPos %s", cmd.PlayerId)
		_, lines := quake3_rcon.SplitReadInfos(res)
		pos, ok := gotoshared.ParsePlayerPosition(lines)
		if !ok {
			logrus.Warnf("Could not parse PlayerPosition from: %v", lines)
			return
		}
		logrus.Debugf("Player position: %+v", pos)

		mapname := cmd.Context.GetCurrentMap()
		existingNames, err := cmd.Context.DB.GetGotoNames(mapname)
		if err != nil {
			logrus.Errorf("GetGotoNames error: %v", err)
			return
		}
		jumpname := gotoshared.ResolveJumpName(existingNames, cmd.Params[0])

		if err := cmd.Context.DB.SaveGoto(mapname, jumpname, pos.PosX, pos.PosY, pos.PosZ, pos.AngleV, pos.AngleH); err != nil {
			logrus.Errorf("SaveGoto error: %v", err)
			return
		}
		logrus.Debugf("Goto saved: map=%s jump=%s", mapname, jumpname)
		cmd.RconText("^5%s^3 saved: %+v", jumpname, pos)
	}
}
