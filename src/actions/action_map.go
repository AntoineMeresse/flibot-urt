package actions

import actionslist "github.com/AntoineMeresse/flibot-urt/src/actions/actions_list"

var Actions = map[string]interface{}{
	"ClientBegin:":           actionslist.ClientBegin,
	"say:":                   actionslist.Say,
	"ClientConnect:":         actionslist.ClientConnect,
	"ClientUserinfo:":        actionslist.ClientUserinfo,
	"Timelimit:":             actionslist.TimelimitHit,
	"ClientUserinfoChanged:": actionslist.DefaultAction,
	"tell:":                  actionslist.Tell,
	"saytell:":               actionslist.SayTell,
}
