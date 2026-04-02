package commandslist

import (
	"fmt"
	"strconv"
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Unban(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		showBanList(cmd)
		return
	}

	if strings.HasPrefix(cmd.Params[0], "@") {
		r, ok := cmd.ResolveAtId(cmd.Params[0])
		if !ok {
			return
		}
		if err := cmd.Context.DB.RemoveBan(r.Id); err != nil {
			cmd.RconText("^1Error removing ban: %s", err.Error())
			return
		}
		cmd.RconText("^7%s^7 has been unbanned.", r.Name)
		return
	}

	id, err := strconv.Atoi(cmd.Params[0])
	if err != nil {
		cmd.RconText("^1Invalid id: %s", cmd.Params[0])
		return
	}

	if err := cmd.Context.DB.RemoveBan(id); err != nil {
		cmd.RconText("^1Error removing ban: %s", err.Error())
		return
	}

	cmd.RconText("^7Player ^3%d^7 has been unbanned.", id)
}

func showBanList(cmd *appcontext.CommandsArgs) {
	bans, err := cmd.Context.DB.GetBans()
	if err != nil {
		cmd.RconText("^1Error fetching ban list: %s", err.Error())
		return
	}
	if len(bans) == 0 {
		cmd.RconText("^7No players are currently banned.")
		return
	}
	entries := make([]string, 0, len(bans))
	for _, b := range bans {
		entries = append(entries, fmt.Sprintf("^7[^3%d^7] ^5%s", b.Id, utils.DecolorString(b.Name)))
	}
	cmd.RconText("^7Players banned (%d): %s", len(bans), strings.Join(entries, "^7, "))
}
