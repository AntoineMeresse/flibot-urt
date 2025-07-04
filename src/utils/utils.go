package utils

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/maruel/natural"
)

func CleanEmptyElements(data []string) []string {
	var res []string
	for _, value := range data {
		if value != "" {
			res = append(res, value)
		}
	}
	return res
}

func CleanDuplicateElements(data []string) []string {
	var res []string
	for _, value := range data {
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

	for k := range myMap {
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
	return re.ReplaceAllString(str, "")
}

func IsDigitOnly(str string) bool {
	return regexp.MustCompile(`^[0-9]+$`).MatchString(str)
}

func ExtractNumber(str string) (int, error) {
	return strconv.Atoi(str)
}

func IsVoteCommand(text string) bool {
	return text == "+" || text == "-" || text == "v"
}

func ToShorterChunkArraySep(strList []string, sep string, exceptFirst bool) []string {
	maxLength := 75
	var res []string
	line := ""

	for i, current := range strList {
		newLine := ""
		if i == 0 && exceptFirst {
			newLine = line + current
		} else {
			newLine = line + current + sep
		}
		if len(newLine) <= maxLength {
			line = newLine
		} else {
			res = append(res, line)
			line = current + sep
		}
	}
	res = append(res, line[0:len(line)-len(sep)])
	return res
}

func ToShorterChunkArray(strList []string) []string {
	return ToShorterChunkArraySep(strList, " ", false)
}

func ToShorterChunkString(str string) []string {
	return ToShorterChunkArray(strings.Split(str, " "))
}

func GetColorRun(i int) string {
	if i == 0 {
		return yellow
	} else if i == 1 {
		return green
	} else if i == 2 {
		return bronze
	}
	return white
}

func truncate(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return math.Round(num*output) / output
}

func RandomFloat(min float64, max float64, precision int) float64 {
	return truncate(rand.Float64()*max+min, precision)
}

func FormatTimeToDate(t time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d", t.Year(), t.Month(), t.Day())
}

func GetTodayDateFormated() string {
	return FormatTimeToDate(time.Now().Local())
}

func BytesNumberConverter(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

func isZero(v string) bool {
	return v == "0" || v == "00.000" || v == "0s"
}

func IsImprovement(v string) bool {
	return !strings.Contains(v, "-") && !isZero(v)
}
