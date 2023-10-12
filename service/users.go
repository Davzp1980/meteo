package service

import (
	"database/sql"
	"meteo/repository"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(db *sql.DB, username, password string) (int, error) {
	hashedPassword, _ := HashePassword(password)

	id, err := repository.CreateUser(db, username, hashedPassword)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func HashePassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
