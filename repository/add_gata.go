package repository

import (
	"database/sql"
	"log"
)

func AddData(db *sql.DB, object string, temperetura float32, humidity, pressure int) error {

	_, err := db.Query("INSERT INTO data (object, temperetura, humidity, pressure) VALUES ($1,$2,$3,$4)", object, temperetura, humidity, pressure)
	if err != nil {
		log.Println("Data adding error", err)
	}

	return nil
}
