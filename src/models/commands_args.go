package models

import (
	"fmt"

	"github.com/sirupsen/logrus"
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

func (c *CommandsArgs) RconGlobalText(text string, a ...any) {
	c.Context.RconText(true, "", text, a...)
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
	cmd := fmt.Sprintf(command, a...)
	logrus.Debugf("Rcon command: %s", cmd)
	return c.Context.Rcon.RconCommand(cmd)
}

func (c *CommandsArgs) RconCommandExtractValue(command string, a ...any) string {
	return c.Context.Rcon.RconCommandExtractValue(fmt.Sprintf(command, a...))
}

func (c *CommandsArgs) GetPlayerGuid() (guid string) {
	player, err := c.Context.Players.GetPlayer(c.PlayerId)

	if (err != nil) {
		logrus.Errorf("Couldn't find player with (id: %s). %s", c.PlayerId, err.Error())
	}

	return player.Guid
}