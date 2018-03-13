package main 

import (
  "github.com/gin-gonic/gin"
)

// InitializeRoutes contains all routes for the application 
func InitializeRoutes() *gin.Engine {
	r := gin.Default()
	return r 
}