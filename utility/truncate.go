package utility

import (
	"fmt"

	"example.com/rest-api/db"
)

func TruncateTable(tableName string) error {
	query := fmt.Sprintf("DELETE FROM %s;", tableName)
	_, err := db.DbConnection.Exec(query)

	if err != nil {
		return err
	}

	// _, err = db.DbConnection.Begin() //for new auto increment

	if err != nil {
		return err
	}

	return nil
}
