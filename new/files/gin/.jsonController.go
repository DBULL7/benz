package controllers 

import (
	"github.com/gin-gonic/gin"
	 c "{{ .Path }}"
)

//HelloWorld auto generated controller
func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
