package commandslist

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils/msg"
)

func MapList(cmd *models.CommandsArgs) {
	maps := cmd.Server.GetMapList()
	new := []string{}
	new = append(new, fmt.Sprintf(msg.MAP_LIST, len(maps)))
	if len(maps) > 5 {
		new = append(new, maps[:5]...)
		new = append(new, "...")
	} else {
		new = append(new, maps...)
	}
	cmd.RconList(new)
}