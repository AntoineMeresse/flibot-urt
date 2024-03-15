package models

import (
	"fmt"
)

type CommandsArgs struct {
	Server *Server
	PlayerId string
	Params []string 
	IsGlobal bool 
}

func (c *CommandsArgs) RconText(text string) {
	if c.IsGlobal {
		c.Server.Rcon.RconCommand(fmt.Sprintf("say ^3%s", text))
	} else {
		c.Server.Rcon.RconCommand(fmt.Sprintf("tell %s ^6[PM] ^3%s", c.PlayerId, text))
	}
}

func (c *CommandsArgs) RconList(list []string) {
	for _, text := range list {
		c.RconText(text)
	}
}

func (c *CommandsArgs) RconCommand(command string) (res string) {
	return c.Server.Rcon.RconCommand(command)
}

func (c *CommandsArgs) RconCommandExtractValue(command string) string {
	return c.Server.Rcon.RconCommandExtractValue(command)
}