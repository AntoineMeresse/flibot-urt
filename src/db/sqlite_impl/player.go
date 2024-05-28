package sqlite_impl

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func createDb_Player() string {
	return `
		CREATE TABLE IF NOT EXISTS player (
			id INTEGER PRIMARY KEY NOT NULL, 
			guid TEXT NOT NULL, 
			name TEXT NOT NULL, 
			ip_address TEXT NOT NULL, 
			time_joined DATETIME, 
			aliases TEXT
		);
	`
}

func (db SqliteDB) SaveNewPlayer(name string, guid string, ip_address string) error {
	req, err := db.DB.Prepare("INSERT INTO player(name, guid, ip_address) values (?, ?, ?)")
	if err != nil {
		return fmt.Errorf("addPlayer sqlite req error. %s", err.Error())
	}

	res, err := req.Exec(name, guid, ip_address)
	if err != nil {
		return fmt.Errorf("addPlayer sqlite req exec error. %s", err.Error())
	}

	id, _ := res.LastInsertId()
	log.Infof("New user successfully created with id : %d", id)

	return nil
}

func (db SqliteDB) UpdatePlayer() error {
	return nil
}