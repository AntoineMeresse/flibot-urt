package commandslist

import (
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"

	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Help(cmd *appcontext.CommandsArgs) {
	cmdList := utils.ToShorterChunkString(strings.Join(utils.NaturalSort(cmd.Params), ", "))
	cmd.RconList(cmdList)
}
