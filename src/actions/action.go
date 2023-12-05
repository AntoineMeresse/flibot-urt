package actions

import "fmt"

func HandleAction(workerId int, action string, action_params []string) {
	// fmt.Printf("\n[Worker %d] ", workerId)
	if val, ok := Actions[action]; ok {
		val.(func([]string))(action_params)
	} else {
		fmt.Printf("  ----> Not a known action: %s", action)
	}
}