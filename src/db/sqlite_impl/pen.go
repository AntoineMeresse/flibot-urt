package sqlite_impl

import (
	"fmt"

	mydb "github.com/AntoineMeresse/flibot-urt/src/db"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
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

// type PenData struct {
// 	name string
// 	size float64
// 	date time.Time
// }

func (db SqliteDB) Pen_add(guid string, size float64) error {
	log.Debug("[Pen_add] Start")
	today := utils.GetTodayDateFormated()
	rows, err := db.createQuery("SELECT size FROM pen where guid=\"%s\" and date=\"%s\"", guid, today)

	if err != nil {
		return fmt.Errorf("Pen_add error in query")
	}

	if rows.Next() {
		var size float64
		rows.Scan(&size)
		return fmt.Errorf("already used pen. Size: %.3f", size)
	}

	return db.sqliteCommit("Pen_add", "INSERT INTO pen(guid, date, size) values (?, ?, ?)", guid, today, size)
}

func (db SqliteDB) Pen_PenOfTheDay() ([]mydb.PenData, error) {
	log.Debug("[Pen_PenOfTheDay] Start")
	fetchMany := 50
	res := []mydb.PenData{}
	
	rows, err := db.createQuery("SELECT name, size, date FROM pen LEFT JOIN player on pen.guid = player.guid WHERE date=\"%s\" ORDER BY size DESC", utils.GetTodayDateFormated())
	
	if err != nil {
		log.Errorf("Pen_PenOfTheDay error in query. Err: %s", err.Error())
		return res, fmt.Errorf("Pen_PenOfTheDay error in query")
	}

	defer rows.Close()

	i := 1
	for rows.Next() && i <= fetchMany {
		current := mydb.PenData{}
		rows.Scan(&current.Name, &current.Size, &current.Date)
		res = append(res, current)
		i++
	}
	
	return res, nil
}