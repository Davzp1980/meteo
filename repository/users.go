package repository

import (
	"database/sql"
	"log"
	"meteo/internal"
)

func CreateUser(db *sql.DB, username, hashedPassword string) (int, error) {
	var user internal.User
	err := db.QueryRow("INSERT INTO users(name, password_hash) VALUES ($1,$2) RETURNING id", username, hashedPassword).Scan(
		&user.ID)

	if err != nil {
		log.Println("User creation error", err)
		return 0, err
	}
	return user.ID, nil
}

func GetUser(db *sql.DB, username string) (string, string, error) {

	var user internal.User
	err := db.QueryRow("SELECT name, password_hash FROM users WHERE name=$1", username).Scan(&user.Name, &user.Password)
	if err != nil {
		log.Println("User search error", err)
		return "", "", err
	}

	return user.Name, user.Password, nil
}
