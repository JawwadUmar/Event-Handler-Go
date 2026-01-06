package models

import (
	"errors"

	"example.com/rest-api/db"
)

type Register struct {
	Id      int64
	UserId  int64
	EventId int64
}

func (r *Register) Save() error {

	query := `INSERT INTO registerations (user_id, event_id) VALUES (?, ?)`
	stmnt, err := db.DbConnection.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmnt.Exec(r.UserId, r.EventId)

	return err
}

func (r *Register) RemoveRegistration() error {
	query := `DELETE FROM registerations WHERE user_id = ? AND event_id = ?`
	stmnt, err := db.DbConnection.Prepare(query)

	if err != nil {
		return err
	}

	sqlResult, err := stmnt.Exec(r.UserId, r.EventId)

	if err != nil {
		return err
	}

	rowsAffected, err := sqlResult.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("No rows affected :)")
	}

	return err
}

func GetAllRegisterations() ([]Register, error) {
	query := `SELECT * FROM registerations`
	rows, err := db.DbConnection.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var registerations []Register

	for rows.Next() {

		var registration Register
		err = rows.Scan(&registration.Id, &registration.UserId, &registration.EventId)

		if err != nil {
			return nil, err
		}

		registerations = append(registerations, registration)
	}

	return registerations, nil
}
