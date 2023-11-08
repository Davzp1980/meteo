package repository

import (
	"database/sql"
	"log"
)

func MigrateDB(db *sql.DB) error {

	CreateTablesQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id serial PRIMARY KEY,
			name VARCHAR (25) NOT NULL UNIQUE,
			password_hash VARCHAR NOT NULL UNIQUE,
			blocked boolean DEFAULT false
		);
		CREATE TABLE IF NOT EXISTS data (
			id serial PRIMARY KEY,
			object VARCHAR (50) NOT NULL,
			temperetura real NOT NULL,
			humidity integer NOT NULL,
			pressure integer NOT NULL,
			date TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP(0)
		);
`
	_, err := db.Exec(CreateTablesQuery)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
