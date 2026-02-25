package commandslist

import (
	"fmt"
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Lookup(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconUsage()
		return
	}

	search := strings.Join(cmd.Params, " ")
	results, err := cmd.Context.DB.LookupPlayers(search)
	if err != nil {
		cmd.RconText("^1Error during lookup: %s", err.Error())
		return
	}

	if len(results) == 0 {
		cmd.RconText("^7No player found for ^5%s", search)
		return
	}

	cmd.RconText("^7Found ^5%d^7 match(es) for ^5%s^7:", len(results), search)
	for _, r := range results {
		aliasesDisplay := r.Aliases
		if aliasesDisplay == "" {
			aliasesDisplay = "none"
		}
		cmd.RconText("%s", fmt.Sprintf("^3%s ^7[id:^5%d^7 lvl:^5%d^7] ip:^5%s ^7aliases:^6%s", r.Name, r.Id, r.Role, r.Ip, aliasesDisplay))
	}
}
