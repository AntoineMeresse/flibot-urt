package models

import (
	"database/sql"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/AntoineMeresse/flibot-urt/src/api"
	quake3_rcon "github.com/AntoineMeresse/quake3-rcon-go"
)

type Server struct {
	Db *sql.DB
	Rcon quake3_rcon.Rcon
	UrtConfig UrtConfig
	Players Players
	Settings ServerSettings
	Api *api.Api
}

type RconFunction func(format string, a ...any)

func (server *Server) Init() {
	server.UrtConfig.init()
	server.InitSettings()
	server.initPlayers()
	server.initApi()
	
	log.Debugf("-------> Flibot started (/connect %s:%s)\n", server.Rcon.ServerIp, server.Rcon.ServerPort)
}

func (server *Server) initPlayers() {
	server.Players = Players{Mutex: sync.RWMutex{}, List: make(map[string]Player)}
}

func (server *Server) initApi() {
	server.Api = &api.Api{}
	server.Api.Init()
}


