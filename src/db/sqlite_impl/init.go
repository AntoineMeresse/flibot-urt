package sqlite_impl

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	log "github.com/sirupsen/logrus"
)

type SqliteDB struct {
	*sql.DB
}

func InitSqliteDb(dbName string) (SqliteDB, error) {
	db, err := sql.Open("sqlite3", dbName);

	if err != nil {
		log.Fatal(err)
	} else {
		// Add mapoptions table
		// Merge checkpoints/utjruns => runs ?
		initTables := fmt.Sprintf("%s\n%s\n%s\n%s", createDb_Player(), createDb_Checkpoints(), createDb_Pen(), createDb_Admin())
		log.Debugf("Init tables: %s", initTables)

		_, err := db.Exec(initTables)

		if err == nil {
			log.Debugf("[SQL] Tables created or already exist.\n")
		} else {
			log.Fatal(err)
		}
	}

	return SqliteDB{DB: db}, err;
}

func InitSqliteDbDevOnly(dbName string) (SqliteDB, error) {
	db, initError := InitSqliteDb(dbName)

	// Exec some methods
	displaySqliteRequestExec(db.SaveNewPlayer("Fliro", "Flitestguid", "fakeip"))
	displaySqliteRequestExec(db.Admin_add("Flitestguid111", 100))
	displaySqliteRequestExec(db.Admin_add_default("Flitestguid_norole211"))
	displaySqliteRequestExec(db.Admin_update("Flitestguid_norole211", 99))

	return db, initError;
}

func displaySqliteRequestExec(err error) {
	if err != nil {
		log.Error(err.Error())
	} else {
		// log.Debug("No error in sql request")
	}
}

func (db SqliteDB) Close() {
	db.DB.Close()
}