package commandslist

import (
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func Quote(cmd *appcontext.CommandsArgs) {
	quote, err := cmd.Context.DB.GetRandomQuote()
	if err != nil {
		log.Errorf("[Quote] Error: %v", err)
		cmd.RconText("No quotes found. Add some with ^5!addQuote")
		return
	}
	cmd.RconText("^7Quote: ^5%s", quote)
}

func AddQuote(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconUsage()
		return
	}

	quote := strings.Join(cmd.Params, " ")
	err := cmd.Context.DB.SaveQuote(quote)
	if err != nil {
		log.Errorf("[AddQuote] Error: %v", err)
		cmd.RconText("Failed to add quote.")
		return
	}

	cmd.RconText("Quote added successfully!")
}
