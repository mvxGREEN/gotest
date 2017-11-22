package main

import (
  "log"
  "fmt"
  "net/http"
  "github.com/julienschmidt/httprouter"
)

func handleGreeting(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello")
}

func main () {
  router := httprouter
  router.GET("/greeting", handleGreeting)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
