package models

import (
	"fmt"
)

type CommandsArgs struct {
	Context *Context
	PlayerId string
	Params []string 
	IsGlobal bool 
	Usage string
}


func (c *CommandsArgs) RconText(text string, a ...any) {
	c.Context.RconText(c.IsGlobal, c.PlayerId, text, a...)
}

func (c *CommandsArgs) RconBigText(text string, a ...any) {
	c.Context.RconBigText(text, a...)
}

func (c *CommandsArgs) RconUsage() {
	c.RconText("^5Usage^3: %s.", c.Usage)
}

func (c *CommandsArgs) RconUsageWithText(text string, a ...any) {
	additionnalText := fmt.Sprintf(text, a...)
	c.RconText("^5Usage^3: %s. %s", c.Usage, additionnalText)
}

func (c *CommandsArgs) RconList(list []string) {
	for _, text := range list {
		c.RconText(text)
	}
}

func (c *CommandsArgs) RconCommand(command string, a ...any) (res string) {
	return c.Context.Rcon.RconCommand(fmt.Sprintf(command, a...))
}

func (c *CommandsArgs) RconCommandExtractValue(command string, a ...any) string {
	return c.Context.Rcon.RconCommandExtractValue(fmt.Sprintf(command, a...))
}