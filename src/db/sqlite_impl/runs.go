package sqlite_impl

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/sirupsen/logrus"
)

func createDbRuns() string {
	return `
		CREATE TABLE IF NOT EXISTS runs (
			id INTEGER PRIMARY KEY NOT NULL, 
			guid TEXT NOT NULL,
			utj INTEGER NOT NULL,
			mapname TEXT NOT NULL, 
			way TEXT NOT NULL, 
			runtime INTEGER NOT NULL, 
			checkpoints TEXT NOT NULL,
			run_date DATETIME NOT NULL, 
			demopath TEXT
		);
	`
}

func (db SqliteDB) HandleRun(info models.PlayerRunInfo, checkpoints []int) error {
	logrus.Debugf("HandleRun: %v | %v", info, checkpoints)
	return nil
}
