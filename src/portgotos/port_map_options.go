package portgotos

import (
	"encoding/json"
	"os"
	"time"

	"github.com/AntoineMeresse/flibot-urt/src/db"
	log "github.com/sirupsen/logrus"
)

func PortMapOptions(filePath string, database db.DataPersister) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Errorf("PortMapOptions: failed to read file %s: %v", filePath, err)
		return
	}

	var raw map[string][]string
	if err := json.Unmarshal(data, &raw); err != nil {
		log.Errorf("PortMapOptions: failed to parse JSON: %v", err)
		return
	}

	start := time.Now()
	total, errors := 0, 0
	i := 0
	for mapname, options := range raw {
		i++
		if mapname == "reset_options" {
			continue
		}
		log.Infof("PortMapOptions: processing map: %s (%d/%d)", mapname, i, len(raw))

		encoded, err := json.Marshal(options)
		if err != nil {
			log.Errorf("PortMapOptions: failed to marshal options for %s: %v", mapname, err)
			errors++
			continue
		}
		if err := database.SetMapOptions(mapname, string(encoded)); err != nil {
			log.Errorf("PortMapOptions: failed to save options for %s: %v", mapname, err)
			errors++
			continue
		}
		total++
	}

	log.Infof("PortMapOptions: imported %d map options (%d errors) in %s", total, errors, time.Since(start).Round(time.Millisecond))
}
