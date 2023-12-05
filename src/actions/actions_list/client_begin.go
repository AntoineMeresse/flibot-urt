package actionslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func ClientBegin(action_params []string, server models.Server) {
	fmt.Printf("\nClient Begin: %v", action_params)
}