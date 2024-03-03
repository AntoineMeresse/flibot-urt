package models

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/AntoineMeresse/flibot-urt/src/utils"
	quake3_rcon "github.com/AntoineMeresse/quake3-rcon-go"
)

type Server struct {
	Db *sql.DB
	Rcon quake3_rcon.Rcon
	UrtPath UrtPath
	Players []Player
	Mapname string
	Nextmap string
	Maplist []string
}

func (server *Server) Init() {
	server.UrtPath.init()
	server.SetMapList()
	server.initMapName()
	server.initNextMapName()
	
	fmt.Printf("\n-------> Flibot started (/connect %s:%s)\n", server.Rcon.ServerIp, server.Rcon.ServerPort)
}

func (server *Server) SetMapList() {
	res := []string{}
	
	file, err := os.Open(server.UrtPath.DownloadPath)
	
	if err == nil {
		names, err := file.Readdirnames(0)
		if err == nil {
			for _, currentFile := range (names) {
				if (strings.HasSuffix(currentFile, ".pk3")) {
					res = append(res, strings.TrimSuffix(currentFile, ".pk3"))
				}
			}
		}
	}

	defer file.Close()
	
	server.Maplist = res;
	fmt.Println(server.Maplist)
}

func (server *Server) initMapName() {
	server.Mapname = server.Rcon.RconCommandExtractValue("mapname")
	fmt.Printf("Current map is: %s\n", server.Mapname)
}

func (server *Server) initNextMapName() {
	if len(server.Maplist) < 2 {
		server.Nextmap = server.Mapname
	} else {
		nextmap := utils.RandomValueFromSlice(server.Maplist)
		for nextmap != "" && nextmap == server.Mapname {
			nextmap = utils.RandomValueFromSlice(server.Maplist)
		}
		server.Nextmap = nextmap
	}
	fmt.Printf("Nexmap is: %s\n", server.Nextmap)
}

func (server Server) RconText(text string, isGlobal bool, playerNumber string) {
	if isGlobal {
		server.Rcon.RconCommand(fmt.Sprintf("say ^3%s", text))
	} else {
		server.Rcon.RconCommand(fmt.Sprintf("tell %s ^6[PM] ^3%s", playerNumber, text))
	}
}

func (server Server) RconList(list []string, isGlobal bool, playerNumber string) {
	for _, text := range list {
		server.RconText(text, isGlobal, playerNumber)
	}
}
