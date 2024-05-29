package sqlite_impl

func createDb_Pen() string{
	return `
		CREATE TABLE IF NOT EXISTS pen (
			id INTEGER PRIMARY KEY NOT NULL, 
			guid TEXT NOT NULL, 
			date DATETIME NOT NULL, 
			size REAL NOT NULL)
		;
	`
}