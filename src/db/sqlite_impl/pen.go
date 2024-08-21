package sqlite_impl

import (
	"fmt"

	mydb "github.com/AntoineMeresse/flibot-urt/src/db"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	log "github.com/sirupsen/logrus"
)

const (
	JOIN_TYPE= "LEFT JOIN"
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

func (db SqliteDB) Pen_PenOfTheDay() (string, []mydb.PenData, error) {
	log.Debug("[Pen_PenOfTheDay] Start")
	today := utils.GetTodayDateFormated()
	datas, err := db.getPenDatas(50, "SELECT name, size, date FROM pen %s player on pen.guid = player.guid WHERE date=\"%s\" ORDER BY size DESC", JOIN_TYPE, today)
	return today, datas, err
}

func (db SqliteDB) Pen_PenHallOfFame() ([]mydb.PenData, error) {
	log.Debug("[Pen_PenHallOfFame] Start")
	return db.getPenDatas(10, "SELECT name, size, date FROM pen %s player on pen.guid = player.guid ORDER BY size DESC", JOIN_TYPE)
}

func (db SqliteDB) Pen_PenHallOfShame() ([]mydb.PenData, error) {
	log.Debug("[Pen_PenHallOfShame] Start")
	return db.getPenDatas(10, "SELECT name, size, date FROM pen %s player on pen.guid = player.guid ORDER BY size ASC", JOIN_TYPE)
}

func (db SqliteDB) getPenDatas(fetchMany int, format string, args ...any) ([]mydb.PenData, error) {
	res := []mydb.PenData{}
	rows, err := db.createQuery(format, args...)
	
	if err != nil {
		log.Errorf("Pen_PenOfTheDay error in query. Err: %s", err.Error())
		return res, fmt.Errorf("Pen_PenOfTheDay error in query")
	}

	defer rows.Close()

	i := 1
	for rows.Next() && i <= fetchMany {
		current := mydb.PenData{}
		if scanErr := rows.Scan(&current.Name, &current.Size, &current.Date); scanErr != nil {
			log.Error(scanErr.Error())
		}
		res = append(res, current)
		log.Tracef("  |--> %d) %v", i, current)
		i++
	}
	
	return res, nil
}