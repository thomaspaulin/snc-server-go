package main

import (
	"net/http"
	"io"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, req *http.Request) {
	Hello(w, req)
}

func Hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func MatchesHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		matches, err := FetchMatches()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encodeErr := encoder.Encode(matches)
		if encodeErr != nil {
			http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	case "POST":
		decoder := json.NewDecoder(req.Body)
		matches := make([]*Match, 0)
		err := decoder.Decode(matches)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer req.Body.Close()
		for _, m := range matches {
			_, err := m.Save()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func SpecificMatchHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	matchID := vars["matchID"]
	if len(vars) == 0 || matchID == "" {
		http.Error(w, "Missing match ID", http.StatusBadRequest)
	}
	switch req.Method {
	case "GET":
		mID, err := strconv.ParseInt(matchID, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		match, err := FetchMatch(uint32(mID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		err = encoder.Encode(match)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	case "PUT":
		http.Error(w, "Not implemented", http.StatusNotImplemented)
		// todo for updating an entire match
	case "PATCH":
		http.Error(w, "Not implemented", http.StatusNotImplemented)
		// todo for updating a particular aspect of a match
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
