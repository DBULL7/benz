package main


func main() {
	
	app := InitializeRoutes()

	app.Logger.Fatal(app.Start(":3000"))
}