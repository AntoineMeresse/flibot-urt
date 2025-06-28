package sqlite_impl

import (
	"fmt"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/sirupsen/logrus"
	"strconv"
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

	time, err := strconv.Atoi(info.Time)

	if err != nil {
		return err
	}

	rows, err := db.createQuery("SELECT runtime from runs where guid=\"%s\" AND mapname=\"%s\" AND way=\"%s\" AND utj=\"%s\"",
		info.Guid, info.Mapname, info.Way, info.Utj)

	if err != nil {
		return err
	}

	if rows.Next() {
		logrus.Debug("HandleRun: Row found")
		var previousTime int
		if errScan := rows.Scan(&previousTime); errScan != nil {
			return errScan
		}

		timeDiff := previousTime - time
		logrus.Debugf("HandleRun: Time diff: %dms", timeDiff)

		if timeDiff > 0 {

			cps := fmt.Sprintf("%v", checkpoints)
			runDate := utils.GetTodayDateFormated()
			req := "UPDATE runs SET runtime=?, checkpoints=?, run_date=? WHERE guid=? AND utj=?"

			logrus.Debugf("HandleRun: Improvement. Need to update. Req: %s", req)

			err = db.sqliteTransaction("HandleRun Update", req, time, cps, runDate, info.Guid, info.Utj)

			if err != nil {
				return err
			}

			logrus.Debugf("HandleRun: Successful update time: %d for guid: %s", time, info.Guid)
		}
	} else {
		logrus.Debugf("HandleRun: No run found. Create a new entry in db")
		cps := fmt.Sprintf("%v", checkpoints)
		runDate := utils.GetTodayDateFormated()
		err = db.sqliteTransaction("HandleRun Create",
			"INSERT INTO runs (guid, utj, mapname, way, runtime, checkpoints, run_date, demopath) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
			info.Guid, info.Utj, info.Mapname, info.Way, time, cps, runDate, info.Demopath)

		if err != nil {
			return err
		}

		logrus.Debugf("HandleRun: Created new entry in db")
	}

	return nil
}
