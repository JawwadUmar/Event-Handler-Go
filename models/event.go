package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `
		INSERT INTO events(name, description, location, dateTime, user_id)
		VALUES (?, ?, ?, ?, ?)
	`

	preparedStatement, err := db.DbConnection.Prepare(query)

	if err != nil {
		return err
	}

	sqlResult, err := preparedStatement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	defer preparedStatement.Close() //Will be closed once the function ends because of whatever reason

	if err != nil {
		return err
	}

	id, err := sqlResult.LastInsertId()

	if err != nil {
		return err
	}

	e.Id = id

	return nil
}

func GetAllEvents() []Event {
	return events
}
