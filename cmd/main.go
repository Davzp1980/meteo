package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"meteo/cors"
	"meteo/handler"

	"meteo/repository"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("BD is not opend")
	}
	defer db.Close()

	repository.MigrateDB(db)

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
	})

	//router.Use(handler.UserIdentification)

	router.HandleFunc("/sign-in", handler.SignIn(db)).Methods("POST")
	router.HandleFunc("/create-user", handler.CreateUser(db)).Methods("POST")

	router.HandleFunc("/add-data", handler.AddData(db))

	router.HandleFunc("/object", handler.GetDataByObject(db)).Methods("GET")
	router.HandleFunc("/temp", handler.GetDataByTemp(db)).Methods("GET")
	router.HandleFunc("/pres", handler.GetDataByPressure(db)).Methods("GET")
	router.HandleFunc("/hum", handler.GetDataByHumidity(db)).Methods("GET")
	router.HandleFunc("/date", handler.GetDataByDate(db)).Methods("GET")

	corsSettings := cors.CorsSettings()
	handler := corsSettings.Handler(router)

	log.Fatal(http.ListenAndServe(":8000", handler))

}
