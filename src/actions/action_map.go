package actions

import (
	actionslist "github.com/AntoineMeresse/flibot-urt/src/actions/actions_list"
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

type ActionFunc func([]string, *appcontext.AppContext)

var Actions = map[string]ActionFunc{
	"ClientBegin:":           actionslist.ClientBegin,
	"say:":                   actionslist.Say,
	"ClientConnect:":         actionslist.ClientConnect,
	"ClientDisconnect:":      actionslist.ClientDisconnect,
	"ClientUserinfo:":        actionslist.ClientUserinfo,
	"ClientSpawn:":           actionslist.ClientSpawn,
	"Timelimit:":             actionslist.TimelimitHit,
	"ClientUserinfoChanged:": actionslist.DefaultAction,
	"tell:":                  actionslist.Tell,
	"saytell:":               actionslist.SayTell,
	"PlayersDump:":           actionslist.PlayersDump,

	// Runs
	"ClientJumpRunStarted:":    actionslist.ClientJumpRunStarted,
	"ClientJumpRunCanceled:":   actionslist.ClientJumpRunCanceled,
	"ClientJumpRunStopped:":    actionslist.ClientJumpRunStopped,
	"ClientJumpRunCheckpoint:": actionslist.ClientJumpRunCheckpoint,
	"RunLog:":                  actionslist.RunLog,
}
