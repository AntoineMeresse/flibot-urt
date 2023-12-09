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

	// Level: 80
	"setgoto" : {commandslist.SetGoto, 80, ""},
	"removegoto" : {commandslist.RemoveGoto, 80, ""},
}

var CommandsShortcut map[string]string = map[string]string {
	"pl" : "play",
	"sp" : "spec",
	"rmgoto" : "removegoto",
}
