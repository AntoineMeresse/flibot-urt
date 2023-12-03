package actions

import "fmt"

func ClientBegin(otherInfos []string) {
	fmt.Printf("Client Begin: %v", otherInfos)
}