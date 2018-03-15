package controllers

import (
	"net/http"
	"github.com/labstack/echo"
)

type Response struct {
  Response  string `json:"Hello"`
}

// Hello 
func Hello(c echo.Context) error {
	r := &Response{
    Response:  "World",
  }
	return c.JSON(http.StatusOK, r)
}