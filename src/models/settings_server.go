package models

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/AntoineMeresse/flibot-urt/src/utils"
	log "github.com/sirupsen/logrus"
)

type ServerSettings struct {
	Mapname string
	Nextmap string
	Maplist []string
}

func (server *Server) SetMapName(mapName string) {
	server.Settings.Mapname = mapName
}

func (server *Server) SetNextMap(nextMapName string) {
	server.Settings.Nextmap = nextMapName
}

func (server *Server) SetMapList() {
	res := []string{}
	
	file, err := os.Open(server.UrtConfig.DownloadPath)
	
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
	
	server.Settings.Maplist = res;
	log.Println(server.Settings.Maplist)
}

func (server *Server) initMapName() {
	server.Settings.Mapname = server.Rcon.RconCommandExtractValue("mapname")
	log.Debugf("Current map is: %s\n", server.Settings.Mapname)
}

func (server *Server) initNextMapName() {
	if len(server.Settings.Maplist) < 2 {
		server.Settings.Nextmap = server.Settings.Mapname
	} else {
		nextmap := utils.RandomValueFromSlice(server.Settings.Maplist)
		for nextmap != "" && nextmap == server.Settings.Mapname {
			nextmap = utils.RandomValueFromSlice(server.Settings.Maplist)
		}
		server.Settings.Nextmap = nextmap
	}
	log.Debugf("Nexmap is: %s\n", server.Settings.Nextmap)
}

func (server *Server) InitSettings() {
	server.SetMapList()
	server.initMapName()
	server.initNextMapName()
}

func (server *Server) IsMapAlreadyDownloaded(mapname string) bool{
	res := slices.Contains(server.GetMapList(), mapname)
	// log.Debugf("IsMapAlreadyDownloaded (%s): %v", mapname, res)
	return res;
}

func (server *Server) GetMapWithCriteria(searchCriteria string) (*string, error) {
	res := []string{}
	
	for _, m := range(server.GetMapList()) {
		if strings.Contains(strings.ToLower(m), strings.ToLower(searchCriteria)) {
			res = append(res, m)
		}
	}

	if len(res) == 1 {
		return &res[0], nil
	} else {
		if len(res) == 0 {
			return nil, fmt.Errorf("no map found using (^6%s^3)", searchCriteria)
		} else {
			var mapList string;
			if len(res) > 3 {
				mapList = strings.Join(res[:3], ", ")
				mapList += " ^5...^3 "
			} else {
				mapList = strings.Join(res, ", ")
			}
			return nil, fmt.Errorf("multiple maps found [^5%d^3] using (^6%s^3): %s ", len(res), searchCriteria, mapList)
		}
	}
}

////////////////////////////////////////////////////////////////

func (server *Server) GetCurrentMap() string {
	return server.Settings.Mapname
}

func (server *Server) GetNextMap() string {
	return server.Settings.Nextmap
}

func (server *Server) GetMapList() []string {
	return server.Settings.Maplist
}

