package main

import (
	"net/http"
	"log"
	"os"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"github.com/thomaspaulin/snc-server/database"
	"github.com/gorilla/mux"
)

func port() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":4242"
}

func main() {
	var err error
	database.DB, err = sql.Open("sqlite3", "./snc.db")
	handle(err)

	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	r.HandleFunc("/hello", Hello)
	r.HandleFunc("/matches", MatchesHandler)
	r.HandleFunc("/matches/{matchID}", SpecificMatchHandler)

	log.Print("Starting up server on port " + port())
	log.Fatal(http.ListenAndServe(port(), nil))
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}