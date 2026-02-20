package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func PortGotos(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) < 1 {
		cmd.RconText(cmd.Usage)
		return
	}

	sourceMap := cmd.Params[0]
	currentMap := cmd.Context.GetCurrentMap()

	names, err := cmd.Context.DB.GetGotoNames(sourceMap)
	if err != nil {
		log.Errorf("[PortGotos] GetGotoNames error: %v", err)
		cmd.RconText("^1Could not retrieve gotos from ^5%s", sourceMap)
		return
	}

	if len(names) == 0 {
		cmd.RconText("^7No gotos found for ^5%s", sourceMap)
		return
	}

	copied := 0
	for _, jumpname := range names {
		g, ok := cmd.Context.DB.GetGoto(sourceMap, jumpname)
		if !ok {
			log.Warnf("[PortGotos] Could not get goto %s from %s", jumpname, sourceMap)
			continue
		}
		if err := cmd.Context.DB.SaveGoto(currentMap, jumpname, g.PosX, g.PosY, g.PosZ, g.AngleV, g.AngleH); err != nil {
			log.Errorf("[PortGotos] SaveGoto error for %s: %v", jumpname, err)
			continue
		}
		copied++
	}

	cmd.RconText("^7Ported ^5%d^7/^5%d^7 gotos from ^5%s^7 to ^5%s", copied, len(names), sourceMap, currentMap)
}
