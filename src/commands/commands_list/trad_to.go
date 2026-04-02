package commandslist

import (
	"slices"
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func TradTo(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) < 2 {
		cmd.RconUsage()
		return
	}

	translateUrl := cmd.Context.UrtConfig.TranslateUrl
	if translateUrl == "" {
		cmd.RconText("^1Translation service not configured.")
		return
	}

	target := strings.ToLower(cmd.Params[0])
	if !slices.Contains(cmd.Context.UrtConfig.TranslateLangs, target) {
		cmd.RconText("^1Language ^3%s^1 is not supported. Available: ^3%s", target, strings.Join(cmd.Context.UrtConfig.TranslateLangs, "^7, ^3"))
		return
	}

	player, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	if err != nil {
		cmd.RconText("^1Could not find your player info.")
		return
	}

	text := strings.Join(cmd.Params[1:], " ")
	result, err := cmd.Context.Translate(translateUrl, text, target)
	if err != nil {
		cmd.RconText("^1Translation service unavailable.")
		return
	}
	if result.Confidence < MinConfidenceTradTo {
		log.Debugf("[tradto] skipped: confidence too low (%.0f%% < %.0f%%) — guessed lang: %s", result.Confidence, MinConfidenceTradTo, result.Lang)
		return
	}

	cmd.Context.RconText(true, "", "^7[^3%s^7->^3%s^7] ^8%s^7: ^8%s", result.Lang, target, player.Name, result.Translated)
}
