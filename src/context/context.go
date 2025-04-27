package context

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
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
	UrtConfig   models.UrtConfig
	Players     models.Players
	Settings    ServerSettings
	Api         *api.Api
	VoteChannel chan models.Vote
	Runs        models.RunsInfo
}

type RconFunction func(format string, a ...any)

func (context *Context) Init() {
	context.UrtConfig.LoadEnvVariables()

	context.initRcon()
	context.initSettings()
	context.initPlayers()
	context.initRuns()
	context.initApi()
	context.initDb()

	log.Debugf("-------> Flibot started (/connect %s:%s)\n", context.Rcon.ServerIp, context.Rcon.ServerPort)
}

func (context *Context) initPlayers() {
	context.Players = models.Players{Mutex: sync.RWMutex{}, PlayerMap: make(map[string]*models.Player)}
}

func (context *Context) initRuns() {
	context.Runs = models.RunsInfo{RunMutex: sync.RWMutex{}, PlayerRuns: make(map[string]*models.RunPlayerInfo), History: make(map[string][]int)}
}

func (context *Context) initApi() {
	context.Api = &api.Api{
		UjmUrl:         context.UrtConfig.ApiConfig.Url,
		Apikey:         context.UrtConfig.ApiConfig.ApiKey,
		BridgeUrl:      "https://ujm-servers.ovh",
		BridgeLocalUrl: "https://ujm-servers.ovh/local",
		Client:         http.Client{Timeout: time.Second * 2},
		ServerUrl:      context.UrtConfig.ServerConfig.GetServerUrl(),
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
	database, dbErr := sqlite_impl.InitSqliteDbDevOnly("test.db?cache=shared&mode=rwc&_journal_mode=WAL&_synchronous=NORMAL")
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

func (context *Context) NewVote(v models.Vote) {
	context.VoteChannel <- v
}
