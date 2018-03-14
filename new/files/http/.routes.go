package main 

import (
    "github.com/julienschmidt/httprouter"
)

// InitializeRoutes contains all routes for the application
func InitializeRoutes() *httprouter.Router {
		r := httprouter.New()
		
		return r 
}