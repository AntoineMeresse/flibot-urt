package utils

import (
	"math/rand"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/maruel/natural"
)

func CleanEmptyElements(datas []string) []string { 
	var res []string
	for _, value := range(datas) {
		if value != "" {
			res = append(res, value)
		}
	}
	return res
}

func CleanDuplicateElements(datas []string) []string {
	var res []string
	for _, value := range(datas) {
		if !slices.Contains(res, value) {
			res = append(res, value)
		}
	}
	return res
}

func RandomValueFromSlice(list []string) string {
	length := len(list)
	if length == 0 {
		return ""
	} else {
		return list[rand.Intn(length)]
	}
}

func GetKeysSorted(myMap map[string][]string) (sortedKeys []string) {
	keys := make([]string, 0, len(myMap))
 
    for k := range myMap{
        keys = append(keys, k)
    }

    sort.Strings(keys)
	return keys
}

func NaturalSort(stringSlice []string) []string { 
	cpy := append([]string{}, stringSlice...)
	sort.Sort(natural.StringSlice(cpy))
	return cpy
}

func DecolorString(str string) string {
	re := regexp.MustCompile(`\^\d`)
	return re.ReplaceAllString(str, "");
}

func IsDigitOnly(str string) bool {
	return regexp.MustCompile(`^[0-9]+$`).MatchString(str)
}

func ExtractNumber(str string) (int, error) {
	return strconv.Atoi(str)
}

func IsVoteCommand(text string) bool {
	return text == "+" || text == "-"
}

func ToShorterChunkArray(strList []string) []string {
	maxLength := 75
	res := []string{}
	lign := ""
	sep := " "

	for _, current := range strList {
		newLign := lign + current + sep 
		if len(newLign) <= maxLength {
			lign = newLign
		} else {
			res = append(res, lign)
			lign = current + sep
		}
	}
	res = append(res, lign)
	return res
}

func ToShorterChunkString(str string) []string {
	return ToShorterChunkArray(strings.Split(str, " "))
}