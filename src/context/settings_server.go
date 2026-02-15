package appcontext

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

func (c *AppContext) SetMapName(mapName string) {
	c.Settings.Mapname = mapName
}

func (c *AppContext) SetNextMap(nextMapName string) {
	log.Debugf("[SetNextMap] Changing nextmap from %s to %s", c.GetNextMap(), nextMapName)
	c.RconCommand("g_nextmap %s", nextMapName)
	c.Settings.Nextmap = nextMapName
}

func (c *AppContext) SetMapList() {
	var res []string

	file, err := os.Open(c.UrtConfig.DownloadPath)
	if err == nil {
		names, err := file.Readdirnames(0)
		if err == nil {
			for _, currentFile := range names {
				if strings.HasSuffix(currentFile, ".pk3") {
					res = append(res, strings.TrimSuffix(currentFile, ".pk3"))
				}
			}
		}
	}
	defer file.Close()

	c.Settings.Maplist = res
	log.Println(c.Settings.Maplist)
}

func (c *AppContext) initMapName() {
	c.Settings.Mapname = c.Rcon.RconCommandExtractValue("mapname")
	log.Debugf("Current map is: %s\n", c.Settings.Mapname)
}

func (c *AppContext) initNextMapName() {
	if len(c.Settings.Maplist) < 2 {
		c.Settings.Nextmap = c.Settings.Mapname
	} else {
		nextmap := utils.RandomValueFromSlice(c.Settings.Maplist)
		for nextmap != "" && nextmap == c.Settings.Mapname {
			nextmap = utils.RandomValueFromSlice(c.Settings.Maplist)
		}
		c.Settings.Nextmap = nextmap
	}
	log.Debugf("Nexmap is: %s\n", c.Settings.Nextmap)
}

func (c *AppContext) initSettings() {
	log.Debug("Initializing settings... [Start]")
	c.SetMapList()
	c.initMapName()
	c.initNextMapName()
	log.Debug("Initializing settings... [End]")
}

func (c *AppContext) IsMapAlreadyDownloaded(mapname string) bool {
	res := slices.Contains(c.GetMapList(), mapname)
	// log.Debugf("IsMapAlreadyDownloaded (%s): %v", mapname, res)
	return res
}

func (c *AppContext) GetMapWithCriteria(searchCriteria string) (uniqueMap *string, err error) {
	res := []string{}

	for _, m := range c.GetMapList() {
		if strings.Contains(strings.ToLower(m), strings.ToLower(searchCriteria)) {
			res = append(res, m)
			log.Debugf("Map found with criteria (%s): %s", searchCriteria, m)
		}
	}

	log.Debugf("List of possible maps: %v", res)

	if len(res) == 1 {
		return &res[0], nil
	} else {
		if len(res) == 0 {
			return nil, fmt.Errorf("no map found using (^6%s^3)", searchCriteria)
		} else {
			var mapList string
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

func (c *AppContext) GetCurrentMap() string {
	return c.Settings.Mapname
}

func (c *AppContext) GetNextMap() string {
	return c.Settings.Nextmap
}

func (c *AppContext) GetMapList() []string {
	return c.Settings.Maplist
}
