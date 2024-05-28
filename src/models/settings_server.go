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

func (context *Context) SetMapName(mapName string) {
	context.Settings.Mapname = mapName
}

func (context *Context) SetNextMap(nextMapName string) {
	context.Settings.Nextmap = nextMapName
}

func (context *Context) SetMapList() {
	res := []string{}
	
	file, err := os.Open(context.UrtConfig.DownloadPath)
	
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
	
	context.Settings.Maplist = res;
	log.Println(context.Settings.Maplist)
}

func (context *Context) initMapName() {
	context.Settings.Mapname = context.Rcon.RconCommandExtractValue("mapname")
	log.Debugf("Current map is: %s\n", context.Settings.Mapname)
}

func (context *Context) initNextMapName() {
	if len(context.Settings.Maplist) < 2 {
		context.Settings.Nextmap = context.Settings.Mapname
	} else {
		nextmap := utils.RandomValueFromSlice(context.Settings.Maplist)
		for nextmap != "" && nextmap == context.Settings.Mapname {
			nextmap = utils.RandomValueFromSlice(context.Settings.Maplist)
		}
		context.Settings.Nextmap = nextmap
	}
	log.Debugf("Nexmap is: %s\n", context.Settings.Nextmap)
}

func (context *Context) initSettings() {
	context.SetMapList()
	context.initMapName()
	context.initNextMapName()
}

func (context *Context) IsMapAlreadyDownloaded(mapname string) bool{
	res := slices.Contains(context.GetMapList(), mapname)
	// log.Debugf("IsMapAlreadyDownloaded (%s): %v", mapname, res)
	return res;
}

func (context *Context) GetMapWithCriteria(searchCriteria string) (uniqueMap *string , err error) {
	res := []string{}
	
	for _, m := range(context.GetMapList()) {
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

func (context *Context) GetCurrentMap() string {
	return context.Settings.Mapname
}

func (context *Context) GetNextMap() string {
	return context.Settings.Nextmap
}

func (context *Context) GetMapList() []string {
	return context.Settings.Maplist
}

