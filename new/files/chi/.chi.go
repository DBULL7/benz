package main

import (
	"net/http"
	"os"
)

func main() {
	app := InitializeRoutes()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":" + port, app)
}