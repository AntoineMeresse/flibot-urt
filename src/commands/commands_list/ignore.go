package commandslist

import (
	"fmt"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Ignore(cmd *appcontext.CommandsArgs) {
	callerGuid := cmd.GetPlayerGuid()
	if callerGuid == "" {
		cmd.RconText("^1Could not identify your player.")
		return
	}

	if len(cmd.Params) == 0 {
		showIgnoreList(cmd, callerGuid)
		return
	}

	target, err := cmd.Context.Players.GetPlayer(cmd.Params[0])
	if err != nil {
		cmd.RconText("%s", err.Error())
		return
	}

	if target.Number == cmd.PlayerId {
		cmd.RconText("^7You cannot ignore yourself.")
		return
	}

	if err := cmd.Context.DB.AddIgnore(callerGuid, target.Guid); err != nil {
		cmd.RconText("^1Error saving ignore: %s", err.Error())
		return
	}

	cmd.RconText("^5%s^7 added to your ignore list.", target.Name)
}

func showIgnoreList(cmd *appcontext.CommandsArgs, callerGuid string) {
	players, err := cmd.Context.DB.GetIgnoredPlayers(callerGuid)
	if err != nil {
		cmd.RconText("^1Error fetching ignore list: %s", err.Error())
		return
	}
	if len(players) == 0 {
		cmd.RconText("^7Your ignore list is empty.")
		return
	}
	lines := make([]string, 0, len(players))
	for _, p := range players {
		lines = append(lines, fmt.Sprintf("^7[^3%d^7] ^5%s", p.Id, utils.DecolorString(p.Name)))
	}
	cmd.RconText("^7Ignored players (%d): ^8(use ^3!unignore [id]^8 to remove)", len(players))
	cmd.RconList(lines)
}
