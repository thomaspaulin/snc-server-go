package main

import (
	"net/http"
	"io"
	"encoding/json"
	"strconv"
	"github.com/gocraft/web"
	"log"
)

func Index(w web.ResponseWriter, req *web.Request) {
	Hello(w, req)
}

func Hello(w web.ResponseWriter, req *web.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// todo figure out a way get context in as well calling methods on the struct
func (ctx *Context) GetMatches(w web.ResponseWriter, req *web.Request) {
	matches, err := FetchMatches(ctx.DB)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encodeErr := encoder.Encode(matches)
	if encodeErr != nil {
		log.Println(err.Error())
		http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
	}
	// The JSON encoder seems to write the OK header so we don't need to do it manually
}

func (ctx *Context) CreateMatches(w web.ResponseWriter, req *web.Request) {
	decoder := json.NewDecoder(req.Body)
	matches := make([]*Match, 0)
	err := decoder.Decode(&matches)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer req.Body.Close()
	for _, m := range matches {
		_, err := m.Save(ctx.DB)
		if err != nil {
			log.Println(err.Error())
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
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	match, err := FetchMatch(ctx.DB, uint32(mID))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(match)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ctx *Context) GetTeams(w web.ResponseWriter, req *web.Request) {
	teams, err := FetchTeams(ctx.DB)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encodeErr := encoder.Encode(teams)
	if encodeErr != nil {
		log.Println(err.Error())
		http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
	}
}

func (ctx *Context) GetSpecificTeam(w web.ResponseWriter, req *web.Request) {
	teamID := req.PathParams["teamID"]
	if teamID == "" {
		http.Error(w, "Missing team ID", http.StatusBadRequest)
	}
	tID, err := strconv.ParseInt(teamID, 10, 32)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	team, err := FetchTeamByID(ctx.DB, uint32(tID))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(team)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ctx *Context) GetRinks(w web.ResponseWriter, req *web.Request) {
	rinks, err := FetchRinks(ctx.DB)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encodeErr := encoder.Encode(rinks)
	if encodeErr != nil {
		log.Println(err.Error())
		http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
	}
}

func (ctx *Context) GetSpecificRink(w web.ResponseWriter, req *web.Request) {
	teamID := req.PathParams["rinkID"]
	if teamID == "" {
		http.Error(w, "Missing rink ID", http.StatusBadRequest)
	}
	rID, err := strconv.ParseInt(teamID, 10, 32)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	rink, err := FetchRinkByID(ctx.DB, uint32(rID))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(rink)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}