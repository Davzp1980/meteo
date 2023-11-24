package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"meteo/internal"
	"meteo/repository"
	"net/http"
)

func GetDataByObject(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input internal.Data

		json.NewDecoder(r.Body).Decode(&input)
		object := r.URL.Query().Get("object")
		fmt.Println(object)
		data, err := repository.GetDataByObject(db, object)
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

		data, err := repository.GetDataByTemp(db, input.Object)
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

		data, err := repository.GetDataByPressure(db, input.Object)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(data)

	}
}

func GetDataByHumidity(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input internal.Data

		json.NewDecoder(r.Body).Decode(&input)

		data, err := repository.GetDataByHumidity(db, input.Object)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(data)

	}
}

func GetDataByDate(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input internal.Data

		json.NewDecoder(r.Body).Decode(&input)

		data, err := repository.GetDataByDate(db, input.Object)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(data)

	}
}
