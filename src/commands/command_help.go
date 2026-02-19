package commands

import (
	"fmt"
	"sort"
	"strings"

	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func buildHelpLines(role int) []string {
	reverseAlias := map[string][]string{}
	for alias, cmd := range Alias {
		reverseAlias[cmd] = append(reverseAlias[cmd], alias)
	}
	for cmd := range reverseAlias {
		reverseAlias[cmd] = utils.NaturalSort(reverseAlias[cmd])
	}

	byLevel := map[int][]string{}
	for name, cmd := range Commands {
		if utils.IsVoteCommand(name) {
			continue
		}
		if cmd.Level <= role {
			byLevel[cmd.Level] = append(byLevel[cmd.Level], name)
		}
	}

	levels := make([]int, 0, len(byLevel))
	for lvl := range byLevel {
		levels = append(levels, lvl)
	}
	sort.Ints(levels)

	arrow := "^7  |---> "
	returnLine := "^7  | "
	sep := "^7, "
	maxLength := 75
	separator := "^7" + strings.Repeat("-", maxLength)

	var res []string
	for _, lvl := range levels {
		res = append(res, separator)
		names := utils.NaturalSort(byLevel[lvl])
		res = append(res, fmt.Sprintf("^3Commands ^5(level %d)^3:", lvl))
		line := arrow
		for i, name := range names {
			aliases := reverseAlias[name]
			var entry string
			if len(aliases) > 0 {
				entry = fmt.Sprintf("^7!%s (^6%s^7)", name, strings.Join(aliases, ", "))
			} else {
				entry = fmt.Sprintf("^7!%s", name)
			}
			if i < len(names)-1 {
				entry += sep
			}
			newLine := line + entry
			if len(newLine) > maxLength && line != arrow {
				res = append(res, line)
				line = returnLine + entry
			} else {
				line = newLine
			}
		}
		if line != arrow {
			res = append(res, line)
		}
	}
	return res
}
