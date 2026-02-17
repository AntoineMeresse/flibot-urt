package goto_shared

import (
	"fmt"
	"slices"
	"strings"
	"unicode"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func DoesPositionExist(c *appcontext.AppContext, location string) bool {
	_, err := c.DB.PositionGet(c.GetCurrentMap(), location)
	return err == nil
}

func getGotosList(c *appcontext.AppContext) []string {
	names, _ := c.DB.PositionList(c.GetCurrentMap())
	return utils.NaturalSort(names)
}

func splitPosition(position string) (elem string, value string) {
	elem = ""
	value = ""
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

func GetDisplayLocation(c *appcontext.AppContext) []string {
	var res []string
	gotos := getGotosList(c)
	if len(gotos) == 0 {
		res = append(res, fmt.Sprintf("^5%s ^1doesn't^3 have locations yet.", c.GetCurrentMap()))
	} else {
		maxLength := 75
		arrow := "^7  |---> "
		returnlign := "^7  | "
		sep := ", "

		gotosGroup := groupGotos(gotos)

		res = append(res, fmt.Sprintf("Goto list for ^5%s^7: ", c.GetCurrentMap()))
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

func GetJumpNameForSavePos(c *appcontext.AppContext, jumpName string) string {
	if len(jumpName) == 1 {
		if unicode.IsLetter(rune(jumpName[0])) {
			gotos := getGotosList(c)
			i := 1
			startPos := fmt.Sprintf("%s%d", jumpName, i)
			for slices.Contains(gotos, startPos) {
				i += 1
				startPos = fmt.Sprintf("%s%d", jumpName, i)
			}
			return startPos
		}
	}
	return jumpName
}

func RemovePosition(c *appcontext.AppContext, location string) bool {
	return c.DB.PositionDelete(c.GetCurrentMap(), location)
}
