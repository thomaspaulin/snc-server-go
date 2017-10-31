package main

import (
	"net/http"
	"io"
	"encoding/json"
	"strconv"
	"github.com/gocraft/web"
)

func Index(w web.ResponseWriter, req *web.Request) {
	Hello(w, req)
}

func Hello(w web.ResponseWriter, req *web.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// todo figure out a way get context in as well calling methods on the struct
func (ctx *Context) GetMatches(w web.ResponseWriter, req *web.Request) {
	matches, err := FetchMatches(ctx.database)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encodeErr := encoder.Encode(matches)
	if encodeErr != nil {
		http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
	}
	// The JSON encoder seems to write the OK header so we don't need to do it manually
}

func (ctx *Context) CreateMatches(w web.ResponseWriter, req *web.Request) {
	decoder := json.NewDecoder(req.Body)
	matches := make([]*Match, 0)
	err := decoder.Decode(matches)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer req.Body.Close()
	for _, m := range matches {
		_, err := m.Save(ctx.database)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (ctx *Context) GetSpecificMatches(w web.ResponseWriter, req *web.Request) {
	matchID := req.PathParams["matchID"]
	if matchID == "" {
		http.Error(w, "Missing match ID", http.StatusBadRequest)
	}
	mID, err := strconv.ParseInt(matchID, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	match, err := FetchMatch(ctx.database, uint32(mID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(match)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ctx *Context) GetSpecificTeam(w web.ResponseWriter, req *web.Request) {
	teamID := req.PathParams["teamID"]
	if teamID == "" {
		http.Error(w, "Missing team ID", http.StatusBadRequest)
	}
	tID, err := strconv.ParseInt(teamID, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	team, err := FetchTeamByID(ctx.database, uint32(tID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}