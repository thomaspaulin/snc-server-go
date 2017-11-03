package main

import (
	"net/http"
	"io"
	"encoding/json"
	"strconv"
	"github.com/gocraft/web"
	"log"
	"github.com/thomaspaulin/snc-server-go/snc"
)

func Index(w web.ResponseWriter, req *web.Request) {
	Hello(w, req)
}

func Hello(w web.ResponseWriter, req *web.Request) {
	io.WriteString(w, "Hello, world!\n")
}

//
// ALL THE CALLER TO CHOOSE CREATE OR UPDATE BASED ON THEIR METHOD BUT IF THEY TRY TO CREATE SOMETHING THAT EXISTS USE AN ERROR CODE TO DENOTE THAT
// 404 FOR UPDATE
// TODO redo the functions here to reflect that
// With great power comes great responsibility (and also flexibility in this case)

// todo figure out a way get context in as well calling methods on the struct

//------------------------------------------------------------------------------------------------//
// Matches
//------------------------------------------------------------------------------------------------//
//func (ctx *Context) GetMatches(w web.ResponseWriter, req *web.Request) {
//	matches, err := FetchMatches(ctx.DB)
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//	w.Header().Set("Content-Type", "application/json")
//	encoder := json.NewEncoder(w)
//	encodeErr := encoder.Encode(matches)
//	if encodeErr != nil {
//		log.Println(err.Error())
//		http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
//	}
//	// The JSON encoder seems to write the OK header so we don't need to do it manually
//}
//
//func (ctx *Context) CreateMatches(w web.ResponseWriter, req *web.Request) {
//	decoder := json.NewDecoder(req.Body)
//	matches := make([]*Match, 0)
//	err := decoder.Decode(&matches)
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//	defer req.Body.Close()
//	for _, m := range matches {
//		_, err := m.Save(ctx.DB)
//		if err != nil {
//			log.Println(err.Error())
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//		}
//	}
//}
//
//func (ctx *Context) GetSpecificMatches(w web.ResponseWriter, req *web.Request) {
//	matchID := req.PathParams["matchID"]
//	if matchID == "" {
//		http.Error(w, "Missing match ID", http.StatusBadRequest)
//	}
//	mID, err := strconv.ParseInt(matchID, 10, 32)
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusBadRequest)
//	}
//	match, err := FetchMatch(ctx.DB, uint32(mID))
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//	w.Header().Set("Content-Type", "application/json")
//	encoder := json.NewEncoder(w)
//	err = encoder.Encode(match)
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}

//------------------------------------------------------------------------------------------------//
// Teams
//------------------------------------------------------------------------------------------------//
func (ctx *Context) GetTeams(w web.ResponseWriter, req *web.Request) {
	teams, err := ctx.TeamService.Teams()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encodeErr := encoder.Encode(teams)
	if encodeErr != nil {
		log.Println(err.Error())
		http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
		return
	}
}

func (ctx *Context) GetSpecificTeam(w web.ResponseWriter, req *web.Request) {
	teamID := req.PathParams["teamID"]
	if teamID == "" {
		http.Error(w, "Missing team ID", http.StatusBadRequest)
		return
	}
	tID, err := strconv.ParseInt(teamID, 10, 32)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	team, err := ctx.TeamService.Team(int(tID))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(team)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//------------------------------------------------------------------------------------------------//
// Rinks
//------------------------------------------------------------------------------------------------//
//func (ctx *Context) GetRinks(w web.ResponseWriter, req *web.Request) {
//	rinks, err := FetchRinks(ctx.DB)
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//	w.Header().Set("Content-Type", "application/json")
//	encoder := json.NewEncoder(w)
//	encodeErr := encoder.Encode(rinks)
//	if encodeErr != nil {
//		log.Println(err.Error())
//		http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
//	}
//}
//
//func (ctx *Context) GetSpecificRink(w web.ResponseWriter, req *web.Request) {
//	teamID := req.PathParams["rinkID"]
//	if teamID == "" {
//		http.Error(w, "Missing rink ID", http.StatusBadRequest)
//	}
//	rID, err := strconv.ParseInt(teamID, 10, 32)
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//	rink, err := FetchRinkByID(ctx.DB, uint32(rID))
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//	w.Header().Set("Content-Type", "application/json")
//	encoder := json.NewEncoder(w)
//	err = encoder.Encode(rink)
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}

//------------------------------------------------------------------------------------------------//
// Divisions
//------------------------------------------------------------------------------------------------//
func (ctx *Context) CreateDivision(w web.ResponseWriter, req *web.Request) {
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()
	d := snc.Division{}
	err := decoder.Decode(&d)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = ctx.DivisionService.CreateDivision(&d)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ctx *Context) GetDivisions(w web.ResponseWriter, req *web.Request) {
	divs, err := ctx.DivisionService.Divisions()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encodeErr := encoder.Encode(divs)
	if encodeErr != nil {
		log.Println(encodeErr.Error())
		http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
		return
	}
}

func (ctx *Context) GetSpecificDivision(w web.ResponseWriter, req *web.Request) {
	divisionID := req.PathParams["divisionID"]
	if divisionID == "" {
		http.Error(w, "Missing division ID", http.StatusBadRequest)
	}
	dID, err := strconv.ParseInt(divisionID, 10, 32)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	div, err := ctx.DivisionService.Division(int(dID))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(div)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ctx *Context) UpdateDivision(w web.ResponseWriter, req *web.Request) {
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()
	d := snc.Division{}
	err := decoder.Decode(&d)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = ctx.DivisionService.UpdateDivision(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
