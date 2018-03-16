package controllers

import (
	"net/http"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
)

// Hello
func Hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := ioutil.ReadFile("../dist/index.html")

	if err == nil {
			w.Write(data)
	} else {
			w.WriteHeader(404)
			w.Write([]byte("404 Something went wrong - " + http.StatusText(404)))
	}
}