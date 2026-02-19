package goto_shared

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"unicode"

	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

type PlayerPosition struct {
	PlayerId int     `json:"i"`
	PosX     float64 `json:"x"`
	PosY     float64 `json:"y"`
	PosZ     float64 `json:"z"`
	AngleV   float64 `json:"v"`
	AngleH   float64 `json:"h"`
}

func ParsePlayerPosition(lines []string) (*PlayerPosition, bool) {
	for _, line := range lines {
		commandPrefix := "PlayerPosition: "
		if strings.Contains(line, commandPrefix) {
			jsonStr := strings.Replace(line, commandPrefix, "", 1)
			var pos PlayerPosition
			if err := json.Unmarshal([]byte(jsonStr), &pos); err == nil {
				return &pos, true
			}
		}
	}
	return nil, false
}

func ResolveJumpName(existingNames []string, jumpName string) string {
	if len(jumpName) == 1 && unicode.IsLetter(rune(jumpName[0])) {
		i := 1
		candidate := fmt.Sprintf("%s%d", jumpName, i)
		for slices.Contains(existingNames, candidate) {
			i++
			candidate = fmt.Sprintf("%s%d", jumpName, i)
		}
		return candidate
	}
	return jumpName
}

func splitPosition(position string) (elem string, value string) {
	for i, r := range position {
		if unicode.IsLetter(r) {
			elem += string(r)
		} else {
			return elem, position[i:]
		}
	}
	return elem, value
}

func groupGotos(gotoPositions []string) map[string][]string {
	res := map[string][]string{}
	for _, pos := range gotoPositions {
		elem, value := splitPosition(pos)
		if len(elem) == 0 {
			res["other"] = append(res["other"], value)
		} else if len(value) == 0 {
			res["other"] = append(res["other"], elem)
		} else {
			res[elem] = append(res[elem], pos)
		}
	}
	return res
}

func BuildDisplayLocation(mapname string, gotos []string) []string {
	var res []string
	if len(gotos) == 0 {
		res = append(res, fmt.Sprintf("^5%s ^1doesn't^3 have locations yet.", mapname))
	} else {
		maxLength := 75
		arrow := "^7  |---> "
		returnlign := "^7  | "
		sep := ", "

		gotosGroup := groupGotos(gotos)

		res = append(res, fmt.Sprintf("Goto list for ^5%s^7: ", mapname))
		for _, k := range utils.GetKeysSorted(gotosGroup) {
			lign := fmt.Sprintf("%s ^2%s^7 : ", arrow, k)
			for _, pos := range gotosGroup[k] {
				newElem := pos + sep
				newLign := lign + newElem
				if len(newLign) <= maxLength {
					lign = newLign
				} else {
					res = append(res, lign)
					lign = returnlign + newElem
				}
			}
			res = append(res, strings.TrimSuffix(lign, sep))
		}
	}
	return res
}
