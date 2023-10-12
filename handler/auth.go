package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"meteo/internal"
	"meteo/repository"
	"meteo/service"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var Jwt_key = []byte("meteo_2023")

func SignIn(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input internal.User

		json.NewDecoder(r.Body).Decode(&input)

		//hashedPassword, _ := service.HashePassword(input.Password)

		userName, hashedPassword, err := repository.GetUser(db, input.Name)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !service.CheckPassword(input.Password, hashedPassword) || input.Name != userName {
			w.Write([]byte("Wrong password or user name"))
			return
		}

		expirationTime := time.Now().Add(5 * time.Minute)

		claims := &Claims{
			Username: userName,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		signedToken, err := token.SignedString(Jwt_key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   signedToken,
			Expires: expirationTime,
		})

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("User %s authorizated", userName)))
	}
}
