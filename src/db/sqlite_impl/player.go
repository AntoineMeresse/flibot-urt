package sqlite_impl

import "github.com/AntoineMeresse/flibot-urt/src/utils"

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

func (db SqliteDB) SaveNewPlayer(name string, guid string, ip_address string) error {
	today := utils.GetTodayDateFormated()
	return db.sqliteTransaction("SaveNewPlayer", "INSERT INTO player(name, guid, ip_address, time_joined, aliases) values (?, ?, ?, ?, ?)", 
		name, guid, ip_address, today, name)
}

func (db SqliteDB) UpdatePlayer() error {
	return nil
}