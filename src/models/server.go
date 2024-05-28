package models

import (
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/AntoineMeresse/flibot-urt/src/api"
	"github.com/AntoineMeresse/flibot-urt/src/db"
	"github.com/AntoineMeresse/flibot-urt/src/db/sqlite_impl"
	quake3_rcon "github.com/AntoineMeresse/quake3-rcon-go"
)

type Server struct {
	DB db.DataPersister
	Rcon quake3_rcon.Rcon
	UrtConfig UrtConfig
	Players Players
	Settings ServerSettings
	Api *api.Api
	VoteChannel chan Vote
}

type RconFunction func(format string, a ...any)

func (server *Server) Init() {
	server.UrtConfig.init()
	server.InitSettings()
	server.initPlayers()
	server.initApi()
	server.initDb()
	
	log.Debugf("-------> Flibot started (/connect %s:%s)\n", server.Rcon.ServerIp, server.Rcon.ServerPort)
}

func (server *Server) initPlayers() {
	server.Players = Players{Mutex: sync.RWMutex{}, List: make(map[string]Player)}
}

func (server *Server) initApi() {
	server.Api = &api.Api{}
	server.Api.Init()
}

func (server *Server) initDb() {
	db, dbErr := sqlite_impl.InitSqliteDbDevOnly("test.db")
	// db, dbErr := sqlite_impl.InitSqliteDb("test.db") 

	if dbErr != nil {
		panic("Error trying to instanciate db")
	} 

	server.DB = db;
}


