package models

import (
	"fmt"
)

type CommandsArgs struct {
	Server *Server
	PlayerId string
	Params []string 
	IsGlobal bool 
	Usage string
}


func (c *CommandsArgs) RconText(text string, a ...any) {
	c.Server.RconText(c.IsGlobal, c.PlayerId, text, a...)
}

func (c *CommandsArgs) RconBigText(text string, a ...any) {
	c.Server.RconBigText(text, a...)
}

func (c *CommandsArgs) RconUsage(text string, a ...any) {
	msg := fmt.Sprintf("^5Usage^3: %s", text)
	c.RconText(msg, a...)
}

func (c *CommandsArgs) RconList(list []string) {
	for _, text := range list {
		c.RconText(text)
	}
}

func (c *CommandsArgs) RconCommand(command string, a ...any) (res string) {
	return c.Server.Rcon.RconCommand(fmt.Sprintf(command, a...))
}

func (c *CommandsArgs) RconCommandExtractValue(command string, a ...any) string {
	return c.Server.Rcon.RconCommandExtractValue(fmt.Sprintf(command, a...))
}