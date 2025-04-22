package commandslist

import (
	"fmt"
	"github.com/AntoineMeresse/flibot-urt/src/context"

	"github.com/AntoineMeresse/flibot-urt/src/utils/msg"
)

func MapList(cmd *context.CommandsArgs) {
	maps := cmd.Context.GetMapList()
	var newMapList []string
	newMapList = append(newMapList, fmt.Sprintf(msg.MAP_LIST, len(maps)))
	if len(maps) > 5 {
		newMapList = append(newMapList, maps[:5]...)
		newMapList = append(newMapList, "...")
	} else {
		newMapList = append(newMapList, maps...)
	}
	cmd.RconList(newMapList)
}
