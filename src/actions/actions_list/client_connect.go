package actionslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func ClientConnect(action_params []string, server models.Server) {
	fmt.Printf("\nClient Connect: %v", action_params)
}