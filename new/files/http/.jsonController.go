package controllers

import (
	"net/http"
	"io"
	"github.com/julienschmidt/httprouter"
)

// Hello
func Hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"hello": "world"}`)
}