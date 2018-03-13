package main 

import (
	"github.com/labstack/echo"

)

// InitializeRoutes contains all application routes 
func InitializeRoutes() *echo.Echo {
	r := echo.New()
  return r 	
}