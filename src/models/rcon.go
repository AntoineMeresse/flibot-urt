package models

import "fmt"

func (server *Server) RconText(isGlobal bool, playerId string, text string, a ...any) {
	msg := fmt.Sprintf(text, a...)
	if isGlobal {
		server.Rcon.RconCommand(fmt.Sprintf("say ^3%s", msg))
	} else {
		server.Rcon.RconCommand(fmt.Sprintf("tell %s ^6[PM] ^3%s", playerId, msg))
	}
}