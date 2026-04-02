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

	source, target, ok := parseLangParam(cmd, cmd.Params[0])
	if !ok {
		return
	}

	player, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	if err != nil {
		cmd.RconText("^1Could not find your player info.")
		return
	}

	text := strings.Join(cmd.Params[1:], " ")
	result, err := cmd.Context.Translate(translateUrl, text, source, target)
	if err != nil {
		cmd.RconText("^1Translation service unavailable.")
		return
	}
	if source == "auto" && result.Confidence < MinConfidenceTradTo {
		log.Debugf("[tradto] skipped: confidence too low (%.0f%% < %.0f%%) — guessed lang: %s", result.Confidence, MinConfidenceTradTo, result.Lang)
		return
	}

	cmd.Context.RconText(true, "", "^7[^3%s^7->^3%s^7] ^8%s^7: ^8%s", result.Lang, target, player.Name, result.Translated)
}

// parseLangParam parses either "it" (target only) or "en->it" (source->target).
// Returns (source, target, ok).
func parseLangParam(cmd *appcontext.CommandsArgs, param string) (source, target string, ok bool) {
	langs := cmd.Context.UrtConfig.TranslateLangs
	param = strings.ToLower(param)

	if parts := strings.SplitN(param, "->", 2); len(parts) == 2 {
		src, tgt := parts[0], parts[1]
		if !slices.Contains(langs, src) {
			cmd.RconText("^1Language ^3%s^1 is not supported. Available: ^3%s", src, strings.Join(langs, "^7, ^3"))
			return "", "", false
		}
		if !slices.Contains(langs, tgt) {
			cmd.RconText("^1Language ^3%s^1 is not supported. Available: ^3%s", tgt, strings.Join(langs, "^7, ^3"))
			return "", "", false
		}
		return src, tgt, true
	}

	if !slices.Contains(langs, param) {
		cmd.RconText("^1Language ^3%s^1 is not supported. Available: ^3%s", param, strings.Join(langs, "^7, ^3"))
		return "", "", false
	}
	return "auto", param, true
}
