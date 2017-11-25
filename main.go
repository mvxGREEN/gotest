package main

import (
  "github.com/julienschmidt/httprouter"
  "log"
  "fmt"
  "net/http"
)

func Index(w http.ResponseWriter, r *http.Request,
  _ httprouter.Params) {
  fmt.Fprintf(w, "Home Page\nAmerican Football\nBasketball\nSoccer")

}

func FootballStreams(w http.ResponseWriter, r *http.Request,
  _ httprouter.Params) {
  fmt.Fprintf(w, "American Football Streaming:\n*insert links*")
}

func BasketballStreams(w http.ResponseWriter, r *http.Request,
  _ httprouter.Params) {
  fmt.Fprintf(w, "Basketball Streaming:\n*insert links*")
}

func SoccerStreams(w http.ResponseWriter, r *http.Request,
_ httprouter.Params) {
  fmt.Fprintf(w, "Football (Soccer) Streaming:\n*insert links*")
}

func main () {
  router := httprouter.New()
  router.GET("/", Index)
  router.GET("/football-streams", FootballStreams)
  router.GET("/basketball-streams", BasketballStreams)
  router.GET("/soccer-streams", SoccerStreams)
  log.Fatal(http.ListenAndServe(":8080", router))
}
