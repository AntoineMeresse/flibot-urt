package appcontext

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"sync"

	"github.com/AntoineMeresse/flibot-urt/src/utils"
	log "github.com/sirupsen/logrus"
)

type ServerSettings struct {
	mu      sync.RWMutex
	Mapname string
	Nextmap string
	Maplist []string
}

func (c *AppContext) SetMapName(mapName string) {
	c.Settings.mu.Lock()
	defer c.Settings.mu.Unlock()
	c.Settings.Mapname = mapName
}

func (c *AppContext) SetNextMap(nextMapName string) {
	log.Debugf("[SetNextMap] Changing nextmap from %s to %s", c.GetNextMap(), nextMapName)
	c.RconCommand("g_nextmap %s", nextMapName)
	c.Settings.mu.Lock()
	defer c.Settings.mu.Unlock()
	c.Settings.Nextmap = nextMapName
}

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Errorf("SetMapList: failed to close directory: %v", err)
	}
}

func (c *AppContext) SetMapList() {
	var res []string

	file, err := os.Open(c.UrtConfig.DownloadPath)
	if err == nil {
		defer closeFile(file)
		names, err := file.Readdirnames(0)
		if err == nil {
			for _, currentFile := range names {
				if strings.HasSuffix(currentFile, ".pk3") {
					res = append(res, strings.TrimSuffix(currentFile, ".pk3"))
				}
			}
		}
	}

	c.Settings.mu.Lock()
	c.Settings.Maplist = res
	c.Settings.mu.Unlock()
	log.Debug(res)
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
	log.Debugf("Nextmap is: %s\n", c.Settings.Nextmap)
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
	var res []string

	for _, m := range c.GetMapList() {
		if strings.Contains(strings.ToLower(m), strings.ToLower(searchCriteria)) {
			res = append(res, m)
			log.Debugf("Map found with criteria (%s): %s", searchCriteria, m)
		}
	}

	log.Debugf("List of possible maps: %v", res)

	if len(res) == 1 {
		return &res[0], nil
	} else if len(res) == 0 {
		return nil, fmt.Errorf("no map found using (^6%s^3)", searchCriteria)
	}

	var mapList string
	if len(res) > 3 {
		mapList = strings.Join(res[:3], ", ")
		mapList += " ^5...^3 "
	} else {
		mapList = strings.Join(res, ", ")
	}
	return nil, fmt.Errorf("multiple maps found [^5%d^3] using (^6%s^3): %s ", len(res), searchCriteria, mapList)
}

////////////////////////////////////////////////////////////////

func (c *AppContext) GetCurrentMap() string {
	c.Settings.mu.RLock()
	defer c.Settings.mu.RUnlock()
	return c.Settings.Mapname
}

func (c *AppContext) GetNextMap() string {
	c.Settings.mu.RLock()
	defer c.Settings.mu.RUnlock()
	return c.Settings.Nextmap
}

func (c *AppContext) GetMapList() []string {
	c.Settings.mu.RLock()
	defer c.Settings.mu.RUnlock()
	return c.Settings.Maplist
}
