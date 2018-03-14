package main 

import (
	"github.com/go-chi/chi"	
)

// InitializeRoutes contains all routes for the application
func InitializeRoutes() *chi.Mux {
	r := chi.NewRouter()
	
	return r 
}