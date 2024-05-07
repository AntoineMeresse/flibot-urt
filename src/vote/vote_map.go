package vote

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
	voteslist "github.com/AntoineMeresse/flibot-urt/src/vote/votes_list"
)

type VoteInfo struct {
	function interface{}
	msgFn interface{}
	messageFormat string
}

var votes map[string]VoteInfo = map[string]VoteInfo {
	"map" : {voteslist.MapVote, voteslist.MapMessage,"Changing map to %s"},
	"cycle" : {voteslist.Cyclemap, voteslist.CyclemapMessage, "Cycling to %s"},
	"extend": {voteslist.Extend, voteslist.ExtendMessage, "Entend %d minutes"},
	"reload" : {voteslist.Reload, noFormatting, "Reload"},
	"nextmap" : {voteslist.Nextmap, voteslist.MapMessage, "Changing nextmap to %s"},
}

func noFormatting(server *models.Server, str string, param string) (bool, string) {
	return true, str;
}
