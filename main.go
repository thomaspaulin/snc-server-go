package main

import (
	"net/http"
	"io"
	"log"
	"os"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"github.com/thomaspaulin/snc-server/snc/database"
)

func index(w http.ResponseWriter, req *http.Request) {
	hello(w, req)
}

func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func matches(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {

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