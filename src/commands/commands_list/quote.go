package commandslist

import (
	"math/rand"
	"strconv"
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

const quotePattern = "^5<< ^3%s ^5>> ^6[%d]"


func Quote(cmd *appcontext.CommandsArgs) {
	var id int
	var quote string
	var err error

	if len(cmd.Params) > 0 {
		quoteId, parseErr := strconv.Atoi(cmd.Params[0])
		if parseErr != nil {
			cmd.RconText("^1Invalid id.")
			return
		}
		id, quote, err = cmd.Context.DB.GetQuoteById(quoteId)
		if err != nil {
			log.Errorf("[Quote] Error: %v", err)
			cmd.RconText("Quote ^3[%d]^7 not found.", quoteId)
			return
		}
	} else {
		id, quote, err = cmd.Context.DB.GetRandomQuote()
		if err != nil {
			log.Errorf("[Quote] Error: %v", err)
			cmd.RconText("No quotes found. Add some with ^5!addquote")
			return
		}
	}

	cmd.RconText(quotePattern, quote, id)
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

func highlightMatch(text, search string) string {
	lower := strings.ToLower(text)
	lowerSearch := strings.ToLower(search)
	var result strings.Builder
	i := 0
	for {
		idx := strings.Index(lower[i:], lowerSearch)
		if idx == -1 {
			result.WriteString(text[i:])
			break
		}
		result.WriteString(text[i : i+idx])
		result.WriteString("^1")
		result.WriteString(text[i+idx : i+idx+len(search)])
		result.WriteString("^3")
		i += idx + len(search)
	}
	return result.String()
}

func FindQuote(cmd *appcontext.CommandsArgs) {
	more := false
	
	if len(cmd.Params) == 0 {
		cmd.RconUsage()
		return
	}

	search := strings.Join(cmd.Params, " ")
	quotes, err := cmd.Context.DB.SearchQuotes(search)
	if err != nil {
		log.Errorf("[FindQuote] Error: %v", err)
		cmd.RconText("Error searching quotes.")
		return
	}

	total := len(quotes)
	if total == 0 {
		cmd.RconText("^7No quotes found for ^5%s^7.", search)
		return
	}

	rand.Shuffle(total, func(i, j int) { quotes[i], quotes[j] = quotes[j], quotes[i] })
	if total > 10 {
		quotes = quotes[:10]
		more = true
	}

	cmd.RconText("^7Found ^5%d^7 quote(s) matching ^6(%s)^3", total, search)
	for _, q := range quotes {
		cmd.RconText(quotePattern, highlightMatch(q.Text, search), q.Id)
	}
	
	if more {
		cmd.RconText("^3[...]")
	}
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
