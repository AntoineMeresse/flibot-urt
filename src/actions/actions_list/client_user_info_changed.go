package actionslist

import (
	"strconv"
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func ClientUserinfoChanged(actionParams []string, c *appcontext.AppContext) {
	log.Debugf("ClientUserinfoChanged: %v", actionParams)
	if len(actionParams) < 2 {
		return
	}
	playerNumber := actionParams[0]
	info := splitInfos(strings.Join(actionParams[1:], ""))
	log.Debugf("ClientUserinfoChanged info: %v", info)
	if t, ok := info["t"]; ok {
		if team, err := strconv.Atoi(t); err == nil {
			c.Players.SetTeam(playerNumber, team)
		}
		if t == "3" {
			c.Runs.RunCanceled(playerNumber)
		}
	}
}
