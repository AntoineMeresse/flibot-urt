package commandslist

import (
	"fmt"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func resolveMap(cmd *appcontext.CommandsArgs, search string, indexStr string) (string, bool) {
	mapName, candidates, err := cmd.Context.ResolveMapWithIndex(search, indexStr)
	if err != nil {
		cmd.RconText("%s", err.Error())
		return "", false
	}
	if len(candidates) > 0 {
		showMapCandidates(cmd, candidates, search)
		return "", false
	}
	return mapName, true
}

func showMapCandidates(cmd *appcontext.CommandsArgs, matches []string, search string) {
	entries := make([]string, len(matches))
	for i, m := range matches {
		entries[i] = fmt.Sprintf("^3[^5%d^3] ^7%s", i+1, m)
	}
	cmd.RconText("^3Multiple maps [^5%d^3] for ^6%s^3:", len(matches), search)
	for _, chunk := range utils.ToShorterChunkArray(entries) {
		cmd.RconText(chunk)
	}
}
