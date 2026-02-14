package appcontext

import (
	"fmt"
)

func (c *AppContext) RconText(isGlobal bool, playerId string, text string, a ...any) {
	msg := fmt.Sprintf(text, a...)
	if isGlobal {
		c.Rcon.RconCommand(fmt.Sprintf("say ^3%s", msg))
	} else {
		c.Rcon.RconCommand(fmt.Sprintf("tell %s ^6[PM] ^3%s", playerId, msg))
	}
}

func (c *AppContext) RconBigText(text string, a ...any) {
	msg := fmt.Sprintf(text, a...)
	c.Rcon.RconCommand(fmt.Sprintf("bigtext \"%s\"", msg))
}

func (c *AppContext) RconPrint(text string, a ...any) {
	msg := fmt.Sprintf(text, a...)
	c.Rcon.RconCommand(fmt.Sprintf("\"%s\"", msg))
}

func (c *AppContext) RconCommand(text string, a ...any) {
	c.Rcon.RconCommand(fmt.Sprintf(text, a...))
}
