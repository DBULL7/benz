package main

import (
	"net/http"
	
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	
	InitializeRoutes()

	e.Logger.Fatal(e.Start(":3000"))
}