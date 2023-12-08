package commands

import commandslist "github.com/AntoineMeresse/flibot-urt/src/commands/commands_list"

var Commands map[string]Command = map[string]Command {
	"play" : {commandslist.Play, 0, ""},
	"spec" : {commandslist.Spec, 0, ""},
	"currentmap" : {commandslist.CurrentMap, 0, ""},
	"nextmap" : {commandslist.NextMap, 0, ""},
	"setgoto" : {commandslist.SetGoto, 80, ""},
}

var CommandsShortcut map[string]string = map[string]string {
	"pl" : "play",
	"sp" : "spec",
}
