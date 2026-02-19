package commandslist

import (
	"encoding/json"
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/sirupsen/logrus"
)

func SetMapOptions(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconUsage()
		return
	}
	mapname := cmd.Context.GetCurrentMap()

	raw := strings.Join(cmd.Params, " ")
	optionAliases := map[string]string{
		"fstam": "g_stamina 2",
		"nstam": "g_stamina 1",
		"noob":  "g_overbounces 0",
		"ob":    "g_overbounces 1",
	}

	parts := strings.Split(raw, ",")
	options := make([]string, 0, len(parts))
	for _, p := range parts {
		opt := strings.TrimSpace(p)
		if opt == "" {
			continue
		}
		if expanded, ok := optionAliases[opt]; ok {
			opt = expanded
		}
		options = append(options, opt)
	}

	data, err := json.Marshal(options)
	if err != nil {
		logrus.Errorf("SetMapOptions marshal error: %v", err)
		return
	}
	if err := cmd.Context.DB.SetMapOptions(mapname, string(data)); err != nil {
		logrus.Errorf("SetMapOptions error: %v", err)
		return
	}
	cmd.RconText("^5%s^3 options saved: ^7%s", mapname, strings.Join(options, ", "))
}
