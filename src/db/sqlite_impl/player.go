package sqlite_impl

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
	return db.sqliteCommit("SaveNewPlayer", "INSERT INTO player(name, guid, ip_address) values (?, ?, ?)", name, guid, ip_address)
}

func (db SqliteDB) UpdatePlayer() error {
	return nil
}