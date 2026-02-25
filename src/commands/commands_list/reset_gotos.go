package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func ResetGotos(cmd *appcontext.CommandsArgs) {
	// Parse params: optional map pattern and optional -f flag
	var mapPattern string
	force := false
	for _, p := range cmd.Params {
		if p == "-f" {
			force = true
		} else {
			mapPattern = p
		}
	}

	// Resolve map name
	var mapname string
	if mapPattern == "" {
		mapname = cmd.Context.GetCurrentMap()
	} else {
		m, err := cmd.Context.GetMapWithCriteria(mapPattern)
		if err != nil {
			cmd.RconText(err.Error())
			return
		}
		mapname = *m
	}

	names, err := cmd.Context.DB.GetGotoNames(mapname)
	if err != nil || len(names) == 0 {
		cmd.RconText("^7No gotos found for map ^5%s", mapname)
		return
	}

	if len(names) > 5 && !force {
		cmd.RconText("^3%d^7 gotos will be deleted for map ^5%s^7. Add ^3-f^7 flag to confirm.", len(names), mapname)
		return
	}

	deleted, err := cmd.Context.DB.DeleteAllGotos(mapname)
	if err != nil {
		cmd.RconText("^1Error deleting gotos: %s", err.Error())
		return
	}

	cmd.RconText("^7Deleted ^5%d^7 goto(s) for map ^5%s^7.", deleted, mapname)
}
