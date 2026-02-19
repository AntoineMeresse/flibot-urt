package actionslist

import (
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func InitGame(actionParams []string, c *appcontext.AppContext) {
	cvarString := strings.Join(actionParams, " ")
	parts := strings.Split(cvarString, "\\")
	for i, part := range parts {
		if part == "mapname" && i+1 < len(parts) {
			c.SetMapName(parts[i+1])
			log.Debugf("InitGame: mapname set to %s", parts[i+1])
			return
		}
	}
	log.Warnf("InitGame: could not parse mapname from: %s", cvarString)
}
