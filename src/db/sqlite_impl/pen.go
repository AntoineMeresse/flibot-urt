package sqlite_impl

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func createDb_Pen() string{
	return `
		CREATE TABLE IF NOT EXISTS pen (
			id INTEGER PRIMARY KEY NOT NULL, 
			guid TEXT NOT NULL, 
			date DATETIME NOT NULL, 
			size REAL NOT NULL)
		;
	`
}

type PenData struct {
	name string
	size float64
	date time.Time
}

// TODO: current_date is utc based, change for local timezone

func (db SqliteDB) Pen_add(guid string, size float64) error {
	rows, err := db.Query(fmt.Sprintf("SELECT size FROM pen where guid=\"%s\" and date=current_date", guid))

	if err != nil {
		return fmt.Errorf("Pen_add error in query")
	}

	if rows.Next() {
		var size float64
		rows.Scan(&size)
		return fmt.Errorf("already used pen. Size: %.3f", size)
	}

	return db.sqliteCommit("Pen_add", "INSERT INTO pen(guid, date, size) values (?, date('now'), ?)", guid, size)
}

func (db SqliteDB) Pen_PenOfTheDay() ([]PenData, error) {
	fetchMany := 50
	res := []PenData{}
	
	rows, err := db.Query("SELECT name, size, date FROM pen INNER JOIN player on pen.guid = player.guid WHERE date = current_date ORDER BY size DESC")
	if err != nil {
		log.Error("Pen_PenOfTheDay error in query")
		return res, fmt.Errorf("Pen_PenOfTheDay error in query")
	}

	i := 1
	for rows.Next() && i <= fetchMany {
		current := PenData{}
		rows.Scan(&current.name, &current.size, &current.date)
		res = append(res, current)
		i++
	}

	rows.Close()

	return res, nil
}