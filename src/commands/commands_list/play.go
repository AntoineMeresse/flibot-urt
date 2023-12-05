package commandslist

import (
	"fmt"
)

func Play(playerNumber string, params []string, isGlobal bool) {
	fmt.Printf("\nPlayer (%s): %v", playerNumber, params)
}