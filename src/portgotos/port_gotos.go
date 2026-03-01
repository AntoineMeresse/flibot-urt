package portgotos

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/AntoineMeresse/flibot-urt/src/db"
	log "github.com/sirupsen/logrus"
)

func PortGotos(rootPath string, database db.DataPersister) {
	entries, err := os.ReadDir(rootPath)
	if err != nil {
		log.Errorf("PortGotos: failed to read directory %s: %v", rootPath, err)
		return
	}

	dirs := make([]os.DirEntry, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry)
		}
	}

	start := time.Now()
	total, errors := 0, 0
	for i, entry := range dirs {
		mapname := entry.Name()
		mapDir := filepath.Join(rootPath, mapname)
		log.Infof("PortGotos: processing folder: %s (%d/%d)", mapname, i+1, len(dirs))

		posFiles, err := filepath.Glob(filepath.Join(mapDir, "*.pos"))
		if err != nil {
			log.Errorf("PortGotos: failed to glob %s: %v", mapDir, err)
			continue
		}

		for _, posFile := range posFiles {
			jumpname := strings.TrimSuffix(filepath.Base(posFile), ".pos")
			posX, posY, posZ, angleV, angleH, err := parsePosFile(posFile)
			if err != nil {
				log.Errorf("PortGotos: %v", err)
				errors++
				continue
			}
			if err := database.SaveGoto(mapname, jumpname, posX, posY, posZ, angleV, angleH); err != nil {
				log.Errorf("PortGotos: failed to save goto %s/%s: %v", mapname, jumpname, err)
				errors++
				continue
			}
			total++
		}
	}

	log.Infof("PortGotos: imported %d gotos (%d errors) in %s", total, errors, time.Since(start).Round(time.Millisecond))
}

func parsePosFile(path string) (posX, posY, posZ, angleV, angleH float64, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, 0, 0, 0, 0, fmt.Errorf("failed to read %s: %w", path, err)
	}

	parts := strings.Split(strings.TrimSpace(string(data)), ",")
	if len(parts) < 5 {
		return 0, 0, 0, 0, 0, fmt.Errorf("invalid pos file %s: expected at least 5 fields, got %d", path, len(parts))
	}

	vals := make([]float64, 5)
	for i := range vals {
		vals[i], err = strconv.ParseFloat(strings.TrimSpace(parts[i]), 64)
		if err != nil {
			return 0, 0, 0, 0, 0, fmt.Errorf("invalid value in %s field %d: %w", path, i, err)
		}
	}

	return vals[0], vals[1], vals[2], vals[3], vals[4], nil
}
