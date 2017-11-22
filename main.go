package main

import (
  "github.com/julienschmidt/httprouter"
  "log"
  "fmt"
  "net/http"
)

func handleGreeting(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  fmt.Fprintf(w, "Hello")
}

func main () {
  router := httprouter.New()
  router.GET("/greeting", handleGreeting)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
