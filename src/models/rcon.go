package models

import "fmt"

func (server *Context) RconText(isGlobal bool, playerId string, text string, a ...any) {
	msg := fmt.Sprintf(text, a...)
	if isGlobal {
		server.Rcon.RconCommand(fmt.Sprintf("say ^3%s", msg))
	} else {
		server.Rcon.RconCommand(fmt.Sprintf("tell %s ^6[PM] ^3%s", playerId, msg))
	}
}

func (server *Context) RconBigText(text string, a ...any) {
	msg := fmt.Sprintf(text, a...)
	server.Rcon.RconCommand(fmt.Sprintf("bigtext \"%s\"", msg))
}

func (server *Context) RconPrint(text string, a ...any) {
	msg := fmt.Sprintf(text, a...)
	server.Rcon.RconCommand(fmt.Sprintf("\"%s\"", msg))
}

func (server *Context) RconCommand(text string, a ...any) {
	server.Rcon.RconCommand(fmt.Sprintf(text, a...))
}