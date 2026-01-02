package models

import (
	"example.com/rest-api/db"
	"example.com/rest-api/utility"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func GetAllUsers() ([]User, error) {
	query := `SELECT * FROM users`

	rows, err := db.DbConnection.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (user *User) Save() error {
	query := `
		INSERT INTO users (email, password)
		VALUES(?, ?)
	`

	stmnt, err := db.DbConnection.Prepare(query)

	if err != nil {
		return err
	}

	defer stmnt.Close()

	hashedPassword, err := utility.HashPassword(user.Password)

	if err != nil {
		return nil
	}

	sqlResult, err := stmnt.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}

	id, err := sqlResult.LastInsertId()

	if err != nil {
		return err
	}

	user.Id = id

	return nil

}
