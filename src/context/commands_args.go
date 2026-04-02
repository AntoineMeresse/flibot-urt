package appcontext

import (
	"fmt"
	"strconv"
	"strings"

	db "github.com/AntoineMeresse/flibot-urt/src/db"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/sirupsen/logrus"
)

type CommandsArgs struct {
	Context       *AppContext
	PlayerId      string
	Params        []string
	IsGlobal      bool
	Usage         string
	CommandExists func(string) bool
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
		c.RconText("%s", text)
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

	if err != nil {
		logrus.Errorf("Couldn't find player with (id: %s). %s", c.PlayerId, err.Error())
		return ""
	}

	return player.Guid
}

func (c *CommandsArgs) ResolveAtId(param string) (db.LookupResult, bool) {
	id, err := strconv.Atoi(strings.TrimPrefix(param, "@"))
	if err != nil {
		c.RconText("^1Invalid id: %s", strings.TrimPrefix(param, "@"))
		return db.LookupResult{}, false
	}
	r, found := c.Context.DB.GetPlayerById(id)
	if !found {
		c.RconText("^7No player found with id ^5%d", id)
		return db.LookupResult{}, false
	}
	return r, true
}

func (c *CommandsArgs) ResolveAdminTarget(search string) (*models.Player, bool) {
	target, err := c.Context.Players.GetPlayer(search)
	if err != nil {
		c.RconText("%s", err.Error())
		return nil, false
	}

	caller, err := c.Context.Players.GetPlayer(c.PlayerId)
	if err != nil {
		c.RconText("%s", err.Error())
		return nil, false
	}

	if target.Role > caller.Role {
		c.RconText("^1You cannot target a player with higher level than yours!")
		return nil, false
	}

	return target, true
}
