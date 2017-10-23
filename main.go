package main

import (
	"net/http"
	"io"
	"log"
	"os"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"github.com/thomaspaulin/snc-server/database"
	"github.com/thomaspaulin/snc-server/models"
	"encoding/json"
)

func index(w http.ResponseWriter, req *http.Request) {
	hello(w, req)
}

func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

//todo break it up to support match IDs (use gorilla mux https://github.com/gorilla/mux)
func matches(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case "GET":
			matches, err := models.FetchMatches(database.DB)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/json")
			encoder := json.NewEncoder(w);
			encodeErr := encoder.Encode(matches)
			if encodeErr != nil {
				http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
		case "POST":
			// todo remove this error when the create and update methods have been saved
			http.Error(w, "Not implemented", http.StatusNotImplemented)
			decoder := json.NewDecoder(req.Body)
			matches := make([]*models.Match, 0)
			err := decoder.Decode(matches)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			defer req.Body.Close()
			for _, m := range matches {
				err := m.Save(database.DB)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
			w.WriteHeader(http.StatusOK)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

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

	// todo handle match IDs
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/matches", matches)

	log.Print("Starting up server on port " + port())
	log.Fatal(http.ListenAndServe(port(), nil))
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}