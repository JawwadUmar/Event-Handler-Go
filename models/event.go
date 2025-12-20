package models

import (
	"time"

	"example.com/rest-api/db"
	// "github.com/aws/aws-sdk-go/private/protocol/query"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

// var events = []Event{}

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

	e.Id = id //seems kind of useless

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * from events`
	rows, err := db.DbConnection.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var events []Event = []Event{}

	for rows.Next() {
		var event Event
		err = rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
