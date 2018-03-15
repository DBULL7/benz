package controllers

import (
	"net/http"
	"io"
)

// Hello 
func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"hello": "world"}`)
}