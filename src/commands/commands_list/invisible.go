package commandslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Invisible(cmd *models.CommandsArgs) {
	cmd.RconCommand(fmt.Sprintf("invisible %s", cmd.PlayerId))
}