package sqlite_impl

import (
	"fmt"
)

func createDb_Admin() string{
	return `
		CREATE TABLE IF NOT EXISTS admin (
			guid TEXT PRIMARY KEY NOT NULL, 
			role INTEGER DEFAULT 1)
		;
	`
}

func (db SqliteDB) Admin_add(guid string, role int) error {
	req, err := db.DB.Prepare("INSERT INTO admin(guid, role) values (?, ?)")
	if err != nil {
		return fmt.Errorf("Admin_add sqlite req error. %s", err.Error())
	}

	_, err = req.Exec(guid, role)
	if err != nil {
		return fmt.Errorf("Admin_add sqlite req exec error. %s", err.Error())
	}

	return nil
}

func (db SqliteDB) Admin_add_default(guid string) error {
	return db.Admin_add(guid, 1)
}

func (db SqliteDB) Admin_update(guid string, role int) error {
	req, err := db.DB.Prepare("UPDATE admin set role=? where guid=?")
	if err != nil {
		return fmt.Errorf("Admin_update sqlite req error. %s", err.Error())
	}

	_, err = req.Exec(role, guid)
	if err != nil {
		return fmt.Errorf("Admin_update sqlite req exec error. %s", err.Error())
	}

	return nil
}
