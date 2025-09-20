package db

func Migrate() {
	schema := `
	CREATE TABLE IF NOT EXISTS locations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		latitude REAL,
		longitude REAL
	);

	CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		quantity INTEGER NOT NULL,
		location_id INTEGER,
		FOREIGN KEY(location_id) REFERENCES locations(id)
	);
	`

	DB.MustExec(schema)
}
