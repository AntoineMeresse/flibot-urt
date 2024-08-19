package commands

import commandslist "github.com/AntoineMeresse/flibot-urt/src/commands/commands_list"

var Commands map[string]Command = map[string]Command {
	// Level: 0
	"play" : {commandslist.Play, 0, ""},
	"spec" : {commandslist.Spec, 0, ""},
	"currentmap" : {commandslist.CurrentMap, 0, ""},
	"nextmap" : {commandslist.NextMap, 0, ""},
	"stamina" : {commandslist.Stamina, 0, ""},
	"ready" : {commandslist.Ready, 0, "!ready"},
	"goto" : {commandslist.Goto, 0, "!goto"},
	"invisible" : {commandslist.Invisible, 0, "!invisible"},
	"loadonce" : {commandslist.Loadonce, 0, "!loadonce"},
	"map" : {commandslist.MapFn, 0, "!map [search]"},
	"maps" : {commandslist.MapList, 0, "!maps"},
	"callvote" : {commandslist.Callvote, 0, "!callvote tocomplete"},
	"+" : {commandslist.VoteYes, 0, ""},
	"-" : {commandslist.VoteNo, 0, ""},
	"help" : {commandslist.Help, 0, ""},
	"mapinfo" : {commandslist.MapInformation, 0, ""},
	"topruns" : {commandslist.ToprunsInformation, 0, ""},
	"top" : {commandslist.TopInformation, 0, ""},

	"afk": {commandslist.Afk, 20, "!afk [playerId/Name]"},
	
	// Level: 80
	"setgoto" : {commandslist.SetGoto, 0, ""},
	"removegoto" : {commandslist.RemoveGoto, 80, ""},
	"download" : {commandslist.Download, 80, ""},
	"timelimit" : {commandslist.Timelimit, 80, "!timelimit [1-999]"},
	"extend" : {commandslist.Extend, 80, "!extend or !extend [1-999]"},

	// Dev Only
	"players" : {commandslist.PlayersList, 0, "!players"},
	"player" : {commandslist.PlayersGet, 0, "!player"},
}

var CommandsShortcut map[string]string = map[string]string {
	"pl" : "play",
	"sp" : "spec",
	"rmgoto" : "removegoto",
	"invi" : "invisible",
	"l1" : "loadonce",
	"lo" : "loadonce",
	"mapget" : "download",
	"mg" : "download",
	"dl" : "download",
	"cv" : "callvote",
	"current": "currentmap",
	"mapinfos": "mapinfo",
	"mi": "mapinfo",
	"tr" : "topruns",
}