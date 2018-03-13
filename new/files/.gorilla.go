package main 

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


func main() {
		app := mux.NewRouter()
		
		InitializeRoutes(app)

		log.Fatal(http.ListenAndServe(":3000", app))
}