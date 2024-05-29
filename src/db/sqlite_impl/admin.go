package sqlite_impl

func createDb_Admin() string{
	return `
		CREATE TABLE IF NOT EXISTS admin (
			guid TEXT PRIMARY KEY NOT NULL, 
			role INTEGER DEFAULT 1)
		;
	`
}

func (db SqliteDB) Admin_add(guid string, role int) error {
	return db.sqliteCommit("Admin_add", "INSERT INTO admin(guid, role) values (?, ?)", guid, role)
}

func (db SqliteDB) Admin_add_default(guid string) error {
	return db.Admin_add(guid, 1)
}

func (db SqliteDB) Admin_update(guid string, role int) error {
	return db.sqliteCommit("Admin_update", "UPDATE admin set role=? where guid=?", role, guid)
}
