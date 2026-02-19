package actionslist

import (
	"encoding/json"
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func InitGame(actionParams []string, c *appcontext.AppContext) {
	cvarString := strings.Join(actionParams, " ")
	parts := strings.Split(cvarString, "\\")
	for i, part := range parts {
		if part == "mapname" && i+1 < len(parts) {
			mapname := parts[i+1]
			c.SetMapName(mapname)
			log.Debugf("InitGame: mapname set to %s", mapname)
			if raw, ok := c.DB.GetMapOptions(mapname); ok {
				var options []string
				if err := json.Unmarshal([]byte(raw), &options); err != nil {
					log.Errorf("InitGame: failed to parse options for %s: %v", mapname, err)
				} else {
					log.Debugf("InitGame: map options for %s: %v", mapname, options)
					for _, opt := range options {
						c.RconCommand(opt)
					}
				}
			} else {
				log.Debugf("InitGame: no options set for %s", mapname)
			}
			return
		}
	}
	log.Warnf("InitGame: could not parse mapname from: %s", cvarString)
}
