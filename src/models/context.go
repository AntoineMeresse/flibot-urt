package models

import (
	"net/http"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/AntoineMeresse/flibot-urt/src/api"
	"github.com/AntoineMeresse/flibot-urt/src/db"
	"github.com/AntoineMeresse/flibot-urt/src/db/sqlite_impl"
	quake3rcon "github.com/AntoineMeresse/quake3-rcon-go"
)

type Context struct {
	DB          db.DataPersister
	Rcon        quake3rcon.Rcon
	UrtConfig   UrtConfig
	Players     Players
	Settings    ServerSettings
	Api         *api.Api
	VoteChannel chan Vote
	Runs        RunsInfo
}

type RconFunction func(format string, a ...any)

func (context *Context) Init() {
	context.UrtConfig.loadEnvVariables()

	context.initRcon()
	context.initSettings()
	context.initPlayers()
	context.initRuns()
	context.initApi()
	context.initDb()

	log.Debugf("-------> Flibot started (/connect %s:%s)\n", context.Rcon.ServerIp, context.Rcon.ServerPort)
}

func (context *Context) initPlayers() {
	context.Players = Players{Mutex: sync.RWMutex{}, PlayerMap: make(map[string]Player)}
}

func (context *Context) initRuns() {
	context.Runs = RunsInfo{RunMutex: sync.RWMutex{}, PlayerRuns: make(map[string]*RunPlayerInfo), History: make(map[string][]int)}
}

func (context *Context) initApi() {
	context.Api = &api.Api{
		UjmUrl:         context.UrtConfig.ApiConfig.Url,
		Apikey:         context.UrtConfig.ApiConfig.ApiKey,
		BridgeUrl:      "https://ujm-servers.ovh",
		BridgeLocalUrl: "https://ujm-servers.ovh/local",
		Client:         http.Client{Timeout: time.Second * 2},
	}
}

func (context *Context) initRcon() {
	context.Rcon = quake3rcon.Rcon{
		ServerIp:   context.UrtConfig.ServerConfig.Ip,
		ServerPort: context.UrtConfig.ServerConfig.Port,
		Password:   context.UrtConfig.ServerConfig.Password,
	}

	context.Rcon.Connect()
}

func (context *Context) initDb() {
	database, dbErr := sqlite_impl.InitSqliteDbDevOnly("test.db")
	// db, dbErr := sqlite_impl.InitSqliteDb("test.db")

	if dbErr != nil {
		panic("Error trying to instantiate db")
	}

	context.DB = database
}

func (context *Context) MapSync() {
	mapSyncErr := context.Api.MapSync()
	if mapSyncErr != nil {
		log.Errorf("Error while trying to sync map: %s", mapSyncErr.Error())
		context.RconCommand("reloadMaps")
		context.SetMapList()
		context.RconText(true, "", "^7Local map sync")
	} else {
		context.RconText(true, "", "^7Bridge map sync (All servers)")
	}
}
