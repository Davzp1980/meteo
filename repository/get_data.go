package repository

import (
	"database/sql"
	"log"
	"meteo/internal"
)

func GetDataByObject(db *sql.DB, object string) ([]internal.Data, error) {
	dataByObject := []internal.Data{}

	rows, err := db.Query("SELECT * from data WHERE object=$1", object)
	if err != nil {
		log.Println("GetDataByObject", err)
		return dataByObject, err
	}

	for rows.Next() {
		var d internal.Data

		if err = rows.Scan(&d.ID, &d.Object, &d.Temperetura, &d.Humidity, &d.Pressure, &d.Date); err != nil {
			log.Println(err)
			return dataByObject, err
		}
		dataByObject = append(dataByObject, d)
	}

	return dataByObject, nil

}

func GetDataByTemp(db *sql.DB, object string) ([]internal.Data, error) {
	dataByTemp := []internal.Data{}

	rows, err := db.Query("SELECT * from data WHERE object=$1 ORDER BY temperetura DESC", object)
	if err != nil {
		log.Println("GetDataByTemp", err)
		return dataByTemp, err
	}

	for rows.Next() {
		var d internal.Data

		if err = rows.Scan(&d.ID, &d.Object, &d.Temperetura, &d.Humidity, &d.Pressure, &d.Date); err != nil {
			log.Println(err)
			return dataByTemp, err
		}
		dataByTemp = append(dataByTemp, d)
	}

	return dataByTemp, nil

}

func GetDataByPressure(db *sql.DB, object string) ([]internal.Data, error) {
	dataByPressure := []internal.Data{}

	rows, err := db.Query("SELECT * from data WHERE object=$1 ORDER BY pressure", object)
	if err != nil {
		log.Println("GetDataByPressure", err)
		return dataByPressure, err
	}

	for rows.Next() {
		var d internal.Data

		if err = rows.Scan(&d.ID, &d.Object, &d.Temperetura, &d.Humidity, &d.Pressure, &d.Date); err != nil {
			log.Println(err)
			return dataByPressure, err
		}
		dataByPressure = append(dataByPressure, d)
	}

	return dataByPressure, nil

}

func GetDataByHumidity(db *sql.DB, object string) ([]internal.Data, error) {
	dataByHumidity := []internal.Data{}

	rows, err := db.Query("SELECT * from data WHERE object=$1 ORDER BY humidity", object)
	if err != nil {
		log.Println("GetDataByPressure", err)
		return dataByHumidity, err
	}

	for rows.Next() {
		var d internal.Data

		if err = rows.Scan(&d.ID, &d.Object, &d.Temperetura, &d.Humidity, &d.Pressure, &d.Date); err != nil {
			log.Println(err)
			return dataByHumidity, err
		}
		dataByHumidity = append(dataByHumidity, d)
	}

	return dataByHumidity, nil

}

func GetDataByDate(db *sql.DB, object string) ([]internal.Data, error) {
	dataByDate := []internal.Data{}

	rows, err := db.Query("SELECT * from data WHERE object=$1 ORDER BY date", object)
	if err != nil {
		log.Println("GetDataByDate", err)
		return dataByDate, err
	}

	for rows.Next() {
		var d internal.Data

		if err = rows.Scan(&d.ID, &d.Object, &d.Temperetura, &d.Humidity, &d.Pressure, &d.Date); err != nil {
			log.Println(err)
			return dataByDate, err
		}
		dataByDate = append(dataByDate, d)
	}

	return dataByDate, nil

}
