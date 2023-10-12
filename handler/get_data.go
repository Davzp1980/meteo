package handler

import (
	"database/sql"
	"encoding/json"
	"meteo/internal"
	"meteo/repository"
	"net/http"
)

func GetDataByObject(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input internal.Data

		json.NewDecoder(r.Body).Decode(&input)

		data, err := repository.GetDataByObject(db, input.Object)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(data)

	}
}

func GetDataByTemp(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input internal.Data

		json.NewDecoder(r.Body).Decode(&input)

		data, err := repository.GetDataByObject(db, input.Object)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(data)

	}
}

func GetDataByPressure(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input internal.Data

		json.NewDecoder(r.Body).Decode(&input)

		data, err := repository.GetDataByObject(db, input.Object)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(data)

	}
}
