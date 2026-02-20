package appcontext

import (
	"context"
	"net/http"
	"strconv"
	"sync"
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
	c.Runs = models.RunsInfo{RunMutex: sync.RWMutex{}, PlayerRuns: make(map[string]*models.RunPlayerInfo), History: make(map[string][]int)}
}

func (c *AppContext) initApi() {
	c.Api = &api.Api{
		UjmUrl:         c.UrtConfig.ApiConfig.Url,
		Apikey:         c.UrtConfig.ApiConfig.ApiKey,
		BridgeUrl:      "https://ujm-servers.ovh",
		BridgeLocalUrl: "https://ujm-servers.ovh/local",
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
	return &player
}

func (c *AppContext) InitPlayer(playerNumber string, guid string, name string, ip string) *models.Player {
	player, found := c.DB.GetPlayerByGuid(guid)
	if !found {
		id, err := c.DB.SaveNewPlayer(name, guid, ip)
		if err != nil {
			log.Errorf("[InitPlayer] Error saving new player: %v", err)
		}
		player = models.Player{Guid: guid, Name: name, Ip: ip, Id: strconv.Itoa(id)}
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
