package models

import (
	"database/sql"
	"fmt"
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

func (server *Server) RconText(text string, isGlobal bool, playerNumber string) {
	if isGlobal {
		server.Rcon.RconCommand(fmt.Sprintf("say ^3%s", text))
	} else {
		server.Rcon.RconCommand(fmt.Sprintf("tell %s ^6[PM] ^3%s", playerNumber, text))
	}
}

func (server *Server) RconList(list []string, isGlobal bool, playerNumber string) {
	for _, text := range list {
		server.RconText(text, isGlobal, playerNumber)
	}
}


