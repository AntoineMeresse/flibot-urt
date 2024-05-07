package commandslist

import (
	"strings"

	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Help(cmd *models.CommandsArgs) {
	cmdList := utils.ToShorterChunkString(strings.Join(utils.NaturalSort(cmd.Params), ", "));
	cmd.RconList(cmdList);
}