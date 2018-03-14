package main

import (
	"os"
)

func main() {
	
	app := InitializeRoutes()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Logger.Fatal(app.Start(":" + port))
}