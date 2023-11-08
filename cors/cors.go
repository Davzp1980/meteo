package cors

import (
	"net/http"

	"github.com/rs/cors"
)

func CorsSettings() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowedOrigins: []string{
			"http://localhost:3000",
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"Content-Type",
		},
		OptionsPassthrough: true,
		ExposedHeaders:     []string{},
		Debug:              true,
	})
	return c
}
