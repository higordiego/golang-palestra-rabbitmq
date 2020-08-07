package database

import (
	"database/sql"
)

// Connection - get connection database
func Connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@/rabbitMQ")

	if err != nil {
		return nil, err
	}

	return db, nil
}
