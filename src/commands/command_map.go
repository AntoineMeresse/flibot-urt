package commands

import commandslist "github.com/AntoineMeresse/flibot-urt/src/commands/commands_list"

var Commands map[string]Command = map[string]Command {
	"play" : {commandslist.Play, 0, ""},
}

var CommandsShortcut map[string]string = map[string]string {
	"pl" : "play",
}
