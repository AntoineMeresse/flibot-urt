package commandslist

import (
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

const MinConfidence = 70.0
const MinConfidenceTradTo = 40.0

func Trad(cmd *appcontext.CommandsArgs) {
	translateUrl := cmd.Context.UrtConfig.TranslateUrl
	if translateUrl == "" {
		cmd.RconText("^1Translation service not configured.")
		return
	}

	if len(cmd.Params) == 0 {
		enabled := cmd.Context.ToggleTrad(cmd.PlayerId)
		if enabled {
			cmd.RconText("^7Auto-translation: ^2On")
		} else {
			cmd.RconText("^7Auto-translation: ^1Off")
		}
		return
	}

	player, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	if err != nil {
		cmd.RconText("^1Could not find your player info.")
		return
	}

	text := strings.Join(cmd.Params, " ")
	result, translErr := cmd.Context.Translate(translateUrl, text, "en")
	if translErr != nil {
		cmd.RconText("^1Translation service unavailable.")
		return
	}
	if result.Lang == "en" {
		cmd.RconText("^7Already in ^3English^7.")
		return
	}
	cmd.Context.RconText(true, "", "^7[^3%s^7->^3en^7] ^8%s^7: ^8%s", result.Lang, player.Name, result.Translated)
}
