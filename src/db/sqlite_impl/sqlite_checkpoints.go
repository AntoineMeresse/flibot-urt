package sqlite_impl

func createDb_Checkpoints() string{
	return `
		CREATE TABLE IF NOT EXISTS checkpoints (
			id INTEGER PRIMARY KEY NOT NULL, 
			guid TEXT NOT NULL, 
			utj INTEGER NOT NULL, 
			mapname TEXT NOT NULL, 
			way TEXT NOT NULL, 
			runtime INTEGER NOT NULL, 
			checkpoints TEXT NOT NULL
		);
	`
}