package actionslist

import (
	"slices"
	"strings"

	"github.com/AntoineMeresse/flibot-urt/src/commands"
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

const minConfidence = 50.0

func Say(actionParams []string, c *appcontext.AppContext) {
	if len(actionParams) >= 3 {
		commands.HandleCommand(actionParams, c)
		if c.UrtConfig.TranslateUrl != "" {
			go autoTrad(actionParams, c)
		}
	}
}

func autoTrad(actionParams []string, c *appcontext.AppContext) {
	playerNumber := actionParams[0]
	playerName := strings.TrimSuffix(actionParams[1], ":")
	message := strings.Join(actionParams[2:], " ")

	if strings.HasPrefix(message, "!") || !c.IsTradEnabled(playerNumber) {
		return
	}
	result, err := c.Translate(c.UrtConfig.TranslateUrl, message, "en")
	if err != nil {
		log.Errorf("[trad] %v", err)
		return
	}
	if result.Confidence < minConfidence {
		c.RconText(false, playerNumber, "^7[trad] ^1Could not detect language with enough confidence ^7(%.0f%%)", result.Confidence)
		return
	}
	if !slices.Contains(c.UrtConfig.TranslateLangs, result.Lang) {
		c.RconText(false, playerNumber, "^7[trad] ^1Language ^3%s^1 is not supported.", result.Lang)
		return
	}
	if result.Lang == "en" {
		return
	}
	c.RconText(false, playerNumber, "^7[^3%s^7] ^8%s^7: ^8%s", result.Lang, playerName, result.Translated)
}
