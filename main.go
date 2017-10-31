package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/thomaspaulin/snc-server-go/database"
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
	// todo
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s", 
			os.Getenv("SNC_USER"), 
			os.Getenv("SNC_PW"), 
			os.Getenv("SNC_HOST"), 
			os.Getenv("SNC_DB"))
	database.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	r.HandleFunc("/hello", Hello)
	r.HandleFunc("/matches", MatchesHandler)
	r.HandleFunc("/matches/{matchID}", SpecificMatchHandler)

	log.Print("Starting up server on port " + port())
	log.Fatal(http.ListenAndServe(port(), nil))
}
