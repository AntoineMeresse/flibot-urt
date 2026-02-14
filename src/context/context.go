package appcontext

import (
	"context"
	"net/http"
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
	c.UrtConfig.LoadEnvVariables()

	c.initRcon()
	c.initSettings()
	c.initPlayers()
	c.initRuns()
	c.initApi()
	c.initDb()

	log.Debugf("-------> Flibot started (/connect %s:%s)\n", c.Rcon.ServerIp, c.Rcon.ServerPort)
}

func (c *AppContext) initPlayers() {
	c.Players = models.Players{Mutex: sync.RWMutex{}, PlayerMap: make(map[string]*models.Player)}
}

func (c *AppContext) initRuns() {
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
	// database, dbErr := sqlite_impl.InitSqliteDbDevOnly("test.db?cache=shared&mode=rwc&_journal_mode=WAL&_synchronous=NORMAL")
	// db, dbErr := sqlite_impl.InitSqliteDb("test.db")

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
