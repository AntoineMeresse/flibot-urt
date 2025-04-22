package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
	"strings"

	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Help(cmd *context.CommandsArgs) {
	cmdList := utils.ToShorterChunkString(strings.Join(utils.NaturalSort(cmd.Params), ", "))
	cmd.RconList(cmdList)
}
