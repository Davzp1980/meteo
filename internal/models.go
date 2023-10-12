package internal

import "time"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Blocker  bool   `json:"blocked"`
}

type Data struct {
	ID          int       `json:"id"`
	Object      string    `json:"object"`
	Temperetura float32   `json:"temperetura"`
	Humidity    int       `json:"humidity"`
	Pressure    int       `json:"pressure"`
	Date        time.Time `json:"date"`
}
