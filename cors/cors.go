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
			http.MethodOptions,
		},
		AllowedOrigins: []string{
			"http://localhost:3000",
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"Content-Type",
			"Access-Control-Allow-Origin",
			"*",
			"200",
		},
		OptionsPassthrough: true,
		ExposedHeaders:     []string{},
		Debug:              true,
	})
	return c
}
