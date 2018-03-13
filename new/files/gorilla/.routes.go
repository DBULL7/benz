package main

import (
	"github.com/gorilla/mux"
)

// InitializeRoutes contains the application routes
func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()
	return r
}