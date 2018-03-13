package main 

import (
	"log"
	"net/http"
)


func main() {
		app := InitializeRoutes()

		log.Fatal(http.ListenAndServe(":3000", app))
}