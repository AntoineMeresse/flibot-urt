package commandslist

import (
	"strconv"
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func Quote(cmd *appcontext.CommandsArgs) {
	id, quote, err := cmd.Context.DB.GetRandomQuote()
	if err != nil {
		log.Errorf("[Quote] Error: %v", err)
		cmd.RconText("No quotes found. Add some with ^5!addquote")
		return
	}
	cmd.RconText("^7Quote: ^5%s ^3[%d]", quote, id)
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

func RemoveQuote(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconUsage()
		return
	}

	id, err := strconv.Atoi(cmd.Params[0])
	if err != nil {
		cmd.RconText("^1Invalid id.")
		return
	}

	err = cmd.Context.DB.DeleteQuote(id)
	if err != nil {
		log.Errorf("[RemoveQuote] Error: %v", err)
		cmd.RconText("Failed to remove quote.")
		return
	}

	cmd.RconText("Quote ^3[%d]^7 removed.", id)
}
