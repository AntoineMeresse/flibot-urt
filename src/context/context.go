package appcontext

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/quake3_rcon"

	log "github.com/sirupsen/logrus"

	"github.com/AntoineMeresse/flibot-urt/src/api"
	"github.com/AntoineMeresse/flibot-urt/src/db"
	"github.com/AntoineMeresse/flibot-urt/src/db/postgres_impl"
)

type AppContext struct {
	DB          db.DataPersister
	Rcon        quake3_rcon.Rcon
	UrtConfig   models.UrtConfig
	Players     models.Players
	Settings    ServerSettings
	Api         *api.Api
	VoteChannel chan models.Vote
	VoteActive  atomic.Bool
	Runs        models.RunsInfo
}

type RconFunction func(format string, a ...any)

func (c *AppContext) Init() {
	c.UrtConfig.LoadConfig()

	c.initRcon()
	c.initSettings()
	c.initPlayers()
	c.initRuns()
	c.initApi()
	c.initDb()

	log.Debugf("-------> Flibot started (/connect %s:%s)\n", c.Rcon.ServerIp, c.Rcon.ServerPort)
	c.RconText(true, "", "^6 Flibot initialized ^5:)")
}

func (c *AppContext) initPlayers() {
	log.Debug("Initializing players...")
	c.Players = models.Players{Mutex: sync.RWMutex{}, PlayerMap: make(map[string]*models.Player)}
}

func (c *AppContext) initRuns() {
	log.Debug("Initializing runs...")
	c.Runs = models.RunsInfo{RunMutex: sync.RWMutex{}, PlayerRuns: make(map[string]*models.RunPlayerInfo), History: make(map[string][]int), CpEnabled: make(map[string]bool)}
}

func (c *AppContext) initApi() {
	bridgeUrl := c.UrtConfig.ApiConfig.BridgeUrl
	c.Api = &api.Api{
		UjmUrl:         c.UrtConfig.ApiConfig.Url,
		Apikey:         c.UrtConfig.ApiConfig.ApiKey,
		BridgeUrl:      bridgeUrl,
		BridgeLocalUrl: bridgeUrl + "/local",
		BridgeApiKey:   c.UrtConfig.ApiConfig.BridgeApiKey,
		Client:         http.Client{Timeout: time.Second * 2},
		ServerUrl:      c.UrtConfig.ServerConfig.GetServerUrl(),
		DiscordWebhook: c.UrtConfig.ApiConfig.DiscordWebhook,
	}
}

func (c *AppContext) initRcon() {
	c.Rcon = quake3_rcon.Rcon{
		ServerIp:   c.UrtConfig.ServerConfig.Ip,
		ServerPort: c.UrtConfig.ServerConfig.Port,
		Password:   c.UrtConfig.ServerConfig.Password,
	}

	c.Rcon.Connect()
}

func (c *AppContext) initDb() {
	log.Debug("Initializing Db...")
	database, dbErr := postgres_impl.InitPostGresqlDb(context.TODO(), c.UrtConfig.DbUri)

	if dbErr != nil {
		log.Fatalf("Error trying to instantiate db. Err: %v", dbErr)
	}

	c.DB = database
	c.registerServer()
}

func (c *AppContext) registerServer() {
	port, err := strconv.Atoi(c.UrtConfig.ServerConfig.Port)
	if err != nil {
		log.Fatalf("Invalid server port: %v", err)
	}
	if err := c.DB.RegisterServer(c.UrtConfig.ServerConfig.Ip, port, c.UrtConfig.ServerConfig.Password, c.UrtConfig.ApiConfig.ChannelId, c.UrtConfig.ApiConfig.ServerName); err != nil {
		log.Errorf("Failed to register server in db: %v", err)
	} else {
		log.Debugf("Server registered: %s:%d", c.UrtConfig.ServerConfig.Ip, port)
	}
}

func (c *AppContext) MapSync() {
	mapSyncErr := c.Api.MapSync()
	if mapSyncErr != nil {
		log.Errorf("Error while trying to sync map: %s", mapSyncErr.Error())
		c.RconCommand("reloadMaps")
		c.SetMapList()
		c.RconText(true, "", "^7Local map sync")
	} else {
		c.RconText(true, "", "^7Bridge map sync (All servers)")
	}
}

func (c *AppContext) NewVote(v models.Vote) {
	c.VoteChannel <- v
}

func (c *AppContext) GivePenCoin(player models.Player) bool {
	if err := c.DB.PenDecrementAttempts(player.Guid); err != nil {
		log.Errorf("GivePenCoin: %v", err)
		return false
	}
	c.RconText(false, player.Number, "[PM] ^7You earned a ^5pencoin^7!")
	return true
}

func (c *AppContext) notifyIgnoredOnline(playerNumber string, guid string) {
	time.Sleep(1 * time.Second)
	ignoredGuids, err := c.DB.GetIgnoredGuids(guid)
	if err != nil || len(ignoredGuids) == 0 {
		return
	}
	ignoredSet := make(map[string]struct{}, len(ignoredGuids))
	for _, g := range ignoredGuids {
		ignoredSet[g] = struct{}{}
	}
	var toSpoof []string
	c.Players.Mutex.RLock()
	for num, p := range c.Players.PlayerMap {
		if num != playerNumber {
			if _, ignored := ignoredSet[p.Guid]; ignored {
				toSpoof = append(toSpoof, num)
			}
		}
	}
	c.Players.Mutex.RUnlock()
	for _, num := range toSpoof {
		c.RconCommand("spoof %s ignore %s", playerNumber, num)
	}
	if len(toSpoof) > 0 {
		c.RconText(false, playerNumber, "^7%d ignored player(s) on this server. ^8(verify with ^3/ignorelist^8 in console)", len(toSpoof))
	}
}

func (c *AppContext) registerPlayer(playerNumber string, player models.Player) *models.Player {
	c.Players.AddPlayer(playerNumber, &player)
	log.Debugf("Player %s not found. Creating it (%v)", playerNumber, player)

	if player.Role == 0 {
		c.RconText(false, playerNumber, "You can register using: ^5!register")
	} else {
		c.RconText(false, playerNumber,
			"Welcome back on server ^5%s^3 [%s]. ^3This is a ^1TEST SERVER^3 so some features might be ^1BROKEN^3.",
			player.Name, player.Id,
		)
	}
	go c.notifyIgnoredOnline(playerNumber, player.Guid)
	return &player
}

func (c *AppContext) UpdatePlayerAliases(player *models.Player) {
	player.Aliases = updateAliases(player.Aliases, player.Name)
	if err := c.DB.UpdatePlayerOnJoin(player.Guid, player.Name, player.Ip, player.Aliases); err != nil {
		log.Errorf("[UpdatePlayerAliases] Error updating player: %v", err)
	}
}

func updateAliases(current []string, name string) []string {
	// Remove existing occurrence to avoid duplicates
	filtered := current[:0]
	for _, a := range current {
		if a != name {
			filtered = append(filtered, a)
		}
	}
	// Prepend current name
	updated := append([]string{name}, filtered...)
	// Cap at 30
	if len(updated) > 30 {
		updated = updated[:30]
	}
	return updated
}

func (c *AppContext) InitPlayer(playerNumber string, guid string, name string, ip string) *models.Player {
	player, found := c.DB.GetPlayerByGuid(guid)
	if !found {
		id, err := c.DB.SaveNewPlayer(name, guid, ip)
		if err != nil {
			log.Errorf("[InitPlayer] Error saving new player: %v", err)
		}
		player = models.Player{Guid: guid, Name: name, Ip: ip, Id: strconv.Itoa(id), Aliases: []string{name}}
	} else {
		player.Aliases = updateAliases(player.Aliases, name)
		player.Name = name
		player.Ip = ip
		if err := c.DB.UpdatePlayerOnJoin(guid, name, ip, player.Aliases); err != nil {
			log.Errorf("[InitPlayer] Error updating player on join: %v", err)
		}
	}
	return c.registerPlayer(playerNumber, player)
}

func (c *AppContext) InitPlayerFromDump(playerNumber string, dump models.DumpPlayer) *models.Player {
	player, found := c.DB.GetPlayerByGuid(dump.GUID)
	if !found {
		id, err := c.DB.SaveNewPlayer(dump.Name, dump.GUID, "")
		if err != nil {
			log.Errorf("[InitPlayerFromDump] Error saving new player: %v", err)
		}
		player = models.Player{Guid: dump.GUID, Name: dump.Name, Id: strconv.Itoa(id)}
	}
	return c.registerPlayer(playerNumber, player)
}
