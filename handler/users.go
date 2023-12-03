package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"meteo/internal"
	"meteo/service"
	"net/http"
)

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input internal.User

		json.NewDecoder(r.Body).Decode(&input)

		if r.Method == "POST" {
			if input.Name == "" || input.Password == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			id, err := service.CreateUser(db, input.Name, input.Password)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(fmt.Sprintf("User created ID: %v, Name: %s", id, input.Name))
		}
	}
}
