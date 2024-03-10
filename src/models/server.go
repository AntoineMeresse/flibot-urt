package models

import (
	"database/sql"
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"

	quake3_rcon "github.com/AntoineMeresse/quake3-rcon-go"
)

type Server struct {
	Db *sql.DB
	Rcon quake3_rcon.Rcon
	UrtPath UrtPath
	Players Players
	Settings ServerSettings
}

func (server *Server) Init() {
	server.UrtPath.init()
	server.InitSettings()
	server.initPlayers()
	
	log.Debugf("-------> Flibot started (/connect %s:%s)\n", server.Rcon.ServerIp, server.Rcon.ServerPort)
}

func (server *Server) initPlayers() {
	server.Players = Players{Mutex: sync.RWMutex{}, List: make(map[string]Player)}
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


