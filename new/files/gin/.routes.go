package main 

import (
	"github.com/gin-gonic/gin"
	"main/controllers/hello"
)

// InitializeRoutes contains all routes for the application 
func InitializeRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/", HelloWorld)
	return r 
}