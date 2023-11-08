package handler

import (
	"database/sql"
	"encoding/json"
	"meteo/internal"
	"meteo/repository"
	"net/http"
)

func AddData(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input internal.Data

		json.NewDecoder(r.Body).Decode(&input)

		err := repository.AddData(db, input.Object, input.Temperetura, input.Humidity, input.Pressure)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("get ok"))
		json.NewEncoder(w).Encode("Data added")
	}
}
