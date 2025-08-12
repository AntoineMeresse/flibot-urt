package appcontext

import (
	"fmt"
)

func (context *AppContext) RconText(isGlobal bool, playerId string, text string, a ...any) {
	msg := fmt.Sprintf(text, a...)
	if isGlobal {
		context.Rcon.RconCommand(fmt.Sprintf("say ^3%s", msg))
	} else {
		context.Rcon.RconCommand(fmt.Sprintf("tell %s ^6[PM] ^3%s", playerId, msg))
	}
}

func (context *AppContext) RconBigText(text string, a ...any) {
	msg := fmt.Sprintf(text, a...)
	context.Rcon.RconCommand(fmt.Sprintf("bigtext \"%s\"", msg))
}

func (context *AppContext) RconPrint(text string, a ...any) {
	msg := fmt.Sprintf(text, a...)
	context.Rcon.RconCommand(fmt.Sprintf("\"%s\"", msg))
}

func (context *AppContext) RconCommand(text string, a ...any) {
	context.Rcon.RconCommand(fmt.Sprintf(text, a...))
}
