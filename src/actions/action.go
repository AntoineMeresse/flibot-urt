package actions

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func HandleAction(workerId int, action string, action_params []string, server models.Server) {
	// fmt.Printf("\n[Worker %d] ", workerId)
	if val, ok := Actions[action]; ok {
		val.(func([]string, models.Server))(action_params, server)
	} else {
		fmt.Printf("  ----> Not a known action: %s", action)
	}
}