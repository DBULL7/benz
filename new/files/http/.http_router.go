package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
)


func main() {
    app := httprouter.New()

		InitializeRoutes()

    log.Fatal(http.ListenAndServe(":3000", app))
}