package actionslist

import (
	"slices"
	"strings"

	"github.com/AntoineMeresse/flibot-urt/src/commands"
	commandslist "github.com/AntoineMeresse/flibot-urt/src/commands/commands_list"
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)


func Say(actionParams []string, c *appcontext.AppContext) {
	if len(actionParams) >= 3 {
		commands.HandleCommand(actionParams, c)
		if c.UrtConfig.TranslateUrl != "" {
			go autoTrad(actionParams, c)
		}
	}
}

func autoTrad(actionParams []string, c *appcontext.AppContext) {
	message := strings.Join(actionParams[2:], " ")
	if strings.HasPrefix(message, "!") || len(message) <= 3 {
		return
	}
	targets := c.TradEnabledPlayers()
	if len(targets) == 0 {
		return
	}
	result, err := c.Translate(c.UrtConfig.TranslateUrl, message, "en")
	if err != nil {
		log.Errorf("[trad] %v", err)
		return
	}
	if result.Confidence < commandslist.MinConfidence {
		log.Debugf("[trad] skipped: confidence too low (%.0f%% < %.0f%%) — guessed lang: %s", result.Confidence, commandslist.MinConfidence, result.Lang)
		return
	}
	if result.Lang == "en" || !slices.Contains(c.UrtConfig.TranslateLangs, result.Lang) {
		return
	}
	playerName := strings.TrimSuffix(actionParams[1], ":")
	for _, playerNumber := range targets {
		c.RconText(false, playerNumber, "^7[^3%s^7] ^8%s^7: ^8%s", result.Lang, playerName, result.Translated)
	}
}
