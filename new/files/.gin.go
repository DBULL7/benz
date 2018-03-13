package main 

import (
  "github.com/gin-gonic/gin"
)

var app *gin.Engine


func main() {
	app = gin.Default()

	InitializeRoutes()
	
	app.Run(":3000")
}

