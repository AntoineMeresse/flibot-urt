package appcontext

import (
	"fmt"
	"log/slog"
)

type CommandsArgs struct {
	Context  *AppContext
	PlayerId string
	Params   []string
	IsGlobal bool
	Usage    string
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
	additionalText := fmt.Sprintf(text, a...)
	c.RconText("^5Usage^3: %s. %s", c.Usage, additionalText)
}

func (c *CommandsArgs) RconList(list []string) {
	for _, text := range list {
		c.RconText(text)
	}
}

func (c *CommandsArgs) RconCommand(command string, a ...any) (res string) {
	cmd := fmt.Sprintf(command, a...)
	slog.Debug("Rcon command", "cmd", cmd)
	return c.Context.Rcon.RconCommand(cmd)
}

func (c *CommandsArgs) RconCommandExtractValue(command string, a ...any) string {
	return c.Context.Rcon.RconCommandExtractValue(fmt.Sprintf(command, a...))
}

func (c *CommandsArgs) GetPlayerGuid() (guid string) {
	player, err := c.Context.Players.GetPlayer(c.PlayerId)

	if err != nil {
		slog.Error("Couldn't find player", "id", c.PlayerId, "err", err)
		return ""
	}

	return player.Guid
}
