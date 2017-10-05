package main

import (
	"net/http"
	"io"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
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
	db, err := gorm.Open("sqlite3", "snc.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate schema
	// Start with match and Team only
	db.AutoMigrate(&snc.Match{}, &snc.Team{})

	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/matches", matches)

	log.Print("Starting up server on port " + port())
	log.Fatal(http.ListenAndServe(port(), nil))
}
