package sqlite_impl

import (
	"fmt"

	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func createDb_Player() string {
	return `
		CREATE TABLE IF NOT EXISTS player (
			id INTEGER PRIMARY KEY NOT NULL, 
			guid TEXT NOT NULL UNIQUE, 
			name TEXT NOT NULL, 
			ip_address TEXT NOT NULL, 
			time_joined DATETIME, 
			aliases TEXT
		);
	`
}

func (db SqliteDB) SaveNewPlayer(name string, guid string, ipAddress string) (int, error) {
	today := utils.GetTodayDateFormated()
	err := db.sqliteTransaction("SaveNewPlayer",
		"INSERT INTO player(name, guid, ip_address, time_joined, aliases) values (?, ?, ?, ?, ?)",
		name, guid, ipAddress, today, name)
	if err != nil {
		return 0, fmt.Errorf("could not save player in db. Error: %s", err.Error())
	}
	return 0, nil
}

func (db SqliteDB) InitRight(guid string) error {
	err := db.sqliteTransaction("InitRight", "INSERT INTO admin(guid, role) values (?, ?)", guid, 0)
	if err != nil {
		return fmt.Errorf("could not save player right in db. Error: %s", err.Error())
	}
	return nil
}

func (db SqliteDB) UpdatePlayer() error {
	return nil
}
