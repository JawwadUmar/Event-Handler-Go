package db

import (
	"database/sql"
	_ "database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

func Init() {

	var err error

	// DbConnection, err := sql.Open("sqlite3", "myDb")
	DbConnection, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to the database")
	}

	DbConnection.SetMaxOpenConns(10)
	DbConnection.SetMaxIdleConns(5)

	creatTables()

}

func creatTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`

	_, err := DbConnection.Exec(createEventsTable)

	if err != nil {
		panic(err)
	}
}
