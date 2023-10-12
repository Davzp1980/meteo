package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

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
	router.Use(handler.UserIdentification)

	router.HandleFunc("/sign-in", handler.SignIn(db))
	router.HandleFunc("/create-user", handler.CreateUser(db))

	router.HandleFunc("/add-data", handler.AddData(db))

	router.HandleFunc("/object", handler.GetDataByObject(db)).Methods("POST")
	router.HandleFunc("/temp", handler.GetDataByTemp(db)).Methods("POST")
	router.HandleFunc("/pres", handler.GetDataByPressure(db)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))

}
