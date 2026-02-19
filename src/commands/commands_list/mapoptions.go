package commandslist

import (
	"encoding/json"
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/sirupsen/logrus"
)

func MapOptions(cmd *appcontext.CommandsArgs) {
	mapname := cmd.Context.GetCurrentMap()
	raw, ok := cmd.Context.DB.GetMapOptions(mapname)
	if !ok {
		cmd.RconText("^5%s^3 has no options set.", mapname)
		return
	}
	var options []string
	if err := json.Unmarshal([]byte(raw), &options); err != nil {
		logrus.Errorf("MapOptions unmarshal error: %v", err)
		return
	}
	cmd.RconText("^5%s^3 options: ^7%s", mapname, strings.Join(options, ", "))
}
