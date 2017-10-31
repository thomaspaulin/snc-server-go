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
	p := os.Getenv("SNC_SERV_PORT")
	if p != "" {
		return ":" + p
	}
	return ":4242"
}

func main() {
	// todo
	username := os.Getenv("SNC_USER")
	host := os.Getenv("SNC_HOST")
	DBName := os.Getenv("SNC_DB")

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s", username, os.Getenv("SNC_PW"), host, DBName)
	database.DB, _ = sql.Open("postgres", connStr)

	log.Printf("Here's what I'm using to connect to the database:\n" +
		"USER: %s\nHOST: %s\nDATABASE: %s", username, host, DBName)
	err := database.DB.Ping();
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
