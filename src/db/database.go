package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDb(dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbName);

	if err != nil {
		log.Fatal(err)
	} else {
		// Add mapoptions table
		// Merge checkpoints/utjruns => runs ?
		initTables := `
			CREATE TABLE IF NOT EXISTS player (
				id INTEGER PRIMARY KEY NOT NULL, 
				guid TEXT NOT NULL, 
				name TEXT NOT NULL, 
				ip_address TEXT NOT NULL, 
				time_joined DATETIME, 
				aliases TEXT
			);

			CREATE TABLE IF NOT EXISTS utjruns (
				id INTEGER PRIMARY KEY NOT NULL, 
				guid TEXT NOT NULL, 
				run_date DATETIME NOT NULL, 
				runtime INTEGER NOT NULL, 
				mapname TEXT NOT NULL,  
				way TEXT NOT NULL, 
				demopath TEXT
			);

			CREATE TABLE IF NOT EXISTS checkpoints (
				id INTEGER PRIMARY KEY NOT NULL, 
				guid TEXT NOT NULL, 
				utj INTEGER NOT NULL, 
				mapname TEXT NOT NULL, 
				way TEXT NOT NULL, 
				runtime INTEGER NOT NULL, 
				checkpoints TEXT NOT NULL
			);
			
			CREATE TABLE IF NOT EXISTS pen (
				id INTEGER PRIMARY KEY NOT NULL, 
				guid TEXT NOT NULL, 
				date DATETIME NOT NULL, 
				size REAL NOT NULL)
			;
		`

		_, err := db.Exec(initTables)

		if err == nil {
			fmt.Printf("\n[SQL] Tables created or already exist.\n")
		} else {
			log.Fatal(err)
		}
	}

	return db, err;
}