package sqlite_impl

import (
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/sirupsen/logrus"
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

func (db SqliteDB) SaveNewPlayer(name string, guid string, ipAddress string) error {
	today := utils.GetTodayDateFormated()
	// Step 1: register player
	errPlayer := db.sqliteTransaction("SaveNewPlayer - Player",
		"INSERT INTO player(name, guid, ip_address, time_joined, aliases) values (?, ?, ?, ?, ?)",
		name, guid, ipAddress, today, name)
	if errPlayer != nil {
		logrus.Errorf("Save new player error: %s", errPlayer.Error())
		//return errPlayer
	}
	// Step 2: Add Rights
	errRight := db.sqliteTransaction("SaveNewPlayer - Player",
		"INSERT INTO admin(guid, role) values (?, ?)",
		guid, 0)
	if errRight != nil {
		logrus.Errorf("Save new player right error: %s", errRight.Error())
		return errRight
	}
	logrus.Infof("New player was register in db. Player guid: %s", guid)
	return nil
}

func (db SqliteDB) UpdatePlayer() error {
	return nil
}
