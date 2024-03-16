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
	msg := fmt.Sprintf(text, a...)
	if c.IsGlobal {
		c.Server.Rcon.RconCommand(fmt.Sprintf("say ^3%s", msg))
	} else {
		c.Server.Rcon.RconCommand(fmt.Sprintf("tell %s ^6[PM] ^3%s", c.PlayerId, msg))
	}
}

func (c *CommandsArgs) RconBigText(text string, a ...any) {
	msg := fmt.Sprintf(text, a...)
	c.Server.Rcon.RconCommand(fmt.Sprintf("bigtext \"%s\"", msg))
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