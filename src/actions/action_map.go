package actions

import actionslist "github.com/AntoineMeresse/flibot-urt/src/actions/actions_list"

var Actions = map[string]interface{}{
	"ClientBegin:":           actionslist.ClientBegin,
	"say:":                   actionslist.Say,
	"ClientConnect:":         actionslist.ClientConnect,
	"ClientUserinfo:":        actionslist.ClientUserinfo,
	"ClientSpawn:":           actionslist.ClientSpawn,
	"Timelimit:":             actionslist.TimelimitHit,
	"ClientUserinfoChanged:": actionslist.DefaultAction,
	"tell:":                  actionslist.Tell,
	"saytell:":               actionslist.SayTell,
	"PlayersDump:":           actionslist.PLayersDump,

	// Runs
	"ClientJumpRunStarted:":    actionslist.ClientJumpRunStarted,
	"ClientJumpRunCanceled:":   actionslist.ClientJumpRunCanceled,
	"ClientJumpRunStopped:":    actionslist.ClientJumpRunStopped,
	"ClientJumpRunCheckpoint:": actionslist.ClientJumpRunCheckpoint,
	"RunLog:":                  actionslist.RunLog,
}
