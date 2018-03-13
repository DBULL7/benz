package main
import (
  "net/http"
)

func main() {
  InitializeRoutes()
  if err := http.ListenAndServe(":3000", nil); err != nil {
    panic(err)
  }
}