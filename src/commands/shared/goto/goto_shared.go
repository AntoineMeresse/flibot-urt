package goto_shared

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"unicode"

	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func DoesPositionExist(server *models.Server, jumpName string) (exists bool, path string) {
	locationPath := fmt.Sprintf("%s/%s/%s.pos", server.UrtPath.GotosPath, server.GetCurrentMap(), jumpName)
	_, err := os.Stat(locationPath)
	return !os.IsNotExist(err), locationPath
}

func getGotosList(server *models.Server) []string {
	mapPath := fmt.Sprintf("%s/%s", server.UrtPath.GotosPath, server.GetCurrentMap())

	file, err := os.Open(mapPath)
	if err != nil {
		return nil
	}
	
	locations, err := file.Readdirnames(0)
	
	if err != nil {
		return nil
	}

	res := []string{}

	for _, v := range (locations) {
		res = append(res, strings.TrimSuffix(v, ".pos"))
	}

	return utils.NaturalSort(res)
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

func groupGotos(gotoPositions []string) map[string][]string{
	res := map[string][]string{}

	for _, pos := range(gotoPositions) {
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

func GetDisplayLocation(server *models.Server) []string {
	res := []string{}
	gotos := getGotosList(server)
	if len(gotos) == 0 {
		res = append(res, fmt.Sprintf("^5%s ^1doesn't^3 have locations yet.", server.GetCurrentMap()))
	} else {
		maxLength := 75
		arrow := "^7  |---> "
        returnlign := "^7  | "
		sep := ", "

		gotosGroup := groupGotos(gotos)

		res = append(res, fmt.Sprintf("Goto list for ^5%s^7: ", server.GetCurrentMap()))
		for _, k := range utils.GetKeysSorted(gotosGroup) {
			lign := fmt.Sprintf("%s ^2%s^7 : ", arrow, k)
			for _, pos := range gotosGroup[k] {
				newElem := pos + sep
				newLign := lign + newElem
				if len(newLign) <= maxLength {
					lign = newLign
				} else {
					res = append(res, lign)
					lign = returnlign+newElem
				}
			}
			res = append(res, strings.TrimSuffix(lign, sep))
		}
	}
	return res
}

func GetJumpNameForSavePos(server *models.Server, jumpName string) string {
	if len(jumpName) == 1 {
		if unicode.IsLetter(rune(jumpName[0])) {
			gotos := getGotosList(server)
			i := 1
			startPos := fmt.Sprintf("%s%d", jumpName, i)
			for slices.Contains(gotos, startPos) {
				i+=1
				startPos = fmt.Sprintf("%s%d", jumpName, i)
			}
			return startPos
		} 
	}
	return jumpName
}

func RemovePosition(server *models.Server, jumpName string) bool {
	exists, path := DoesPositionExist(server, jumpName)
	err := os.Remove(path)
	if err != nil {
		return false
	}
	return exists
}
