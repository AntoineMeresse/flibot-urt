package actions

import "fmt"

var Actions map[string]interface{} = map[string]interface{}{
	"ClientBegin:": ClientBegin,
	"say:": Say,
}

func HandleAction(workerId int, action string, otherInfos []string) {
	fmt.Printf("\n[Worker %d] ", workerId)
	if val, ok := Actions[action]; ok {
		val.(func([]string))(otherInfos)
	} else {
		fmt.Printf("  ----> Not a known action: %s", action)
	}
}
