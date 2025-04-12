package vote

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
	voteslist "github.com/AntoineMeresse/flibot-urt/src/vote/votes_list"
)

type VoteInfo struct {
	function      any
	msgFn         any
	messageFormat string
	Usage         string
}

var Votes map[string]VoteInfo = map[string]VoteInfo{
	"map":     {voteslist.MapVote, voteslist.MapMessage, "Changing map to %s", "map [mapname]^3"},
	"cycle":   {voteslist.Cyclemap, voteslist.CyclemapMessage, "Cycling to %s", "cycle^3"},
	"extend":  {voteslist.Extend, voteslist.ExtendMessage, "Extend %d minute(s)", "extend [minutes*]^3. Default: ^61h^3"},
	"reload":  {voteslist.Reload, noFormatting, "Reload", "reload"},
	"nextmap": {voteslist.Nextmap, voteslist.MapMessage, "Changing nextmap to %s", "nextmap [mapname]^3"},
}

func noFormatting(context *models.Context, str string, param string) (bool, string) {
	return true, str
}
