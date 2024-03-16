package commandslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func MapList(cmd *models.CommandsArgs) {
	maps := cmd.Server.GetMapList()
	new := []string{}
	new = append(new, fmt.Sprintf("Server map list [^5%d^3]: ", len(maps)))
	if len(maps) > 5 {
		new = append(new, maps[:5]...)
		new = append(new, "...")
	} else {
		new = append(new, maps...)
	}
	cmd.RconList(new)
}