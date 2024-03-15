package commandslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Loadonce(cmd *models.CommandsArgs) {
	cmd.RconCommand(fmt.Sprintf("simpleload %s", cmd.PlayerId))
}