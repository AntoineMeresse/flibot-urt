package bridge

import (
	"time"

	"github.com/AntoineMeresse/flibot-urt/src/api"
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

const interval = 10 * time.Second

func SendServerInfoToBridge(c *appcontext.AppContext) {
	if c.Api.BridgeUrl == "" {
		log.Info("[bridge] bridgeUrl not configured, server info sync disabled")
		return
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		sendServerInfo(c)
	}
}

func sendServerInfo(c *appcontext.AppContext) {
	mapname := c.GetCurrentMap()

	c.Players.Mutex.RLock()
	players := make([]api.BridgePlayer, 0, len(c.Players.PlayerMap))
	for _, p := range c.Players.PlayerMap {
		if p.Ip == "127.0.0.1" {
			continue
		}
		players = append(players, api.BridgePlayer{
			Name:    p.Name,
			Ingame:  !p.IsSpec(),
			Running: c.Runs.IsRunning(p.Number),
		})
	}
	c.Players.Mutex.RUnlock()

	if err := c.Api.SendServerInfo(mapname, players); err != nil {
		log.Debugf("[bridge] SendServerInfo error: %v", err)
	}
}
