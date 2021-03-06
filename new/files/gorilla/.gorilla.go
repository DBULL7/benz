package main 

import (
	"log"
	"net/http"
	"os"
)


func main() {
	app := InitializeRoutes()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(http.ListenAndServe(":" + port, app))
}