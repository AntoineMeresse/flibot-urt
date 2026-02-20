package commandslist

import (
	"encoding/json"
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func PortMapOptions(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) < 1 {
		cmd.RconText(cmd.Usage)
		return
	}

	sourceMap := cmd.Params[0]
	currentMap := cmd.Context.GetCurrentMap()

	raw, found := cmd.Context.DB.GetMapOptions(sourceMap)
	if !found {
		cmd.RconText("^7No options found for ^5%s", sourceMap)
		return
	}

	if err := cmd.Context.DB.SetMapOptions(currentMap, raw); err != nil {
		log.Errorf("[PortMapOptions] SetMapOptions error: %v", err)
		cmd.RconText("^1Failed to save options for ^5%s", currentMap)
		return
	}

	var options []string
	if err := json.Unmarshal([]byte(raw), &options); err != nil {
		log.Errorf("[PortMapOptions] Unmarshal error: %v", err)
		return
	}

	for _, opt := range cmd.Context.UrtConfig.ResetOptions {
		cmd.RconCommand("%s", opt)
	}
	for _, opt := range options {
		cmd.RconCommand("%s", opt)
	}

	cmd.RconText("^7Options from ^5%s^7 ported to ^5%s^7: ^7%s", sourceMap, currentMap, strings.Join(options, ", "))
}
