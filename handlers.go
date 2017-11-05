package main

import (
	"github.com/thomaspaulin/snc-server-go/snc"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"errors"
)

type Services struct {
	DivisionService		snc.DivisionService
	TeamService			snc.TeamService
	RinkService			snc.RinkService
}

func services(c *gin.Context) (*Services, error) {
	s, exists := c.Get("services")
	if !exists {
		return nil, errors.New("handlers: services struct doesn't exist in the context but it should have been set on init")
	}
	return s.(*Services), nil
}

func Index(c *gin.Context) {
	Hello(c)
}

func Hello(c *gin.Context) {
	c.String(200, "Hello, %s!\n", "world")
}

// todo

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
func GetTeams(c *gin.Context) {
	log.Println("all teams")
	s, err := services(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The server couldn't process the request"})
	}
	teams, err := s.TeamService.Teams()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if len(teams) == 0 {
		teams = make([]*snc.Team, 0)
	}
	c.JSON(http.StatusOK, teams)
}

func GetSpecificTeam(c *gin.Context) {
	s, err := services(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The server couldn't process the request"})
	}
	teamID := c.Param("teamID")
	if teamID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing team ID"})
	} else if _, err := strconv.Atoi(teamID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	tID, err := strconv.Atoi(teamID)
	t, err := s.TeamService.Team(int(tID))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, t)
}

//------------------------------------------------------------------------------------------------//
// Rinks
//------------------------------------------------------------------------------------------------//
func CreateRink(c *gin.Context) {
	s, err := services(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The server couldn't process the request"})
	}
	r := snc.Rink{}
	if err = c.BindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if err = s.RinkService.CreateRink(&r); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

func GetRinks(c *gin.Context) {
	s, err := services(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The server couldn't process the request"})
	}
	rinks, err := s.RinkService.Rinks()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if len(rinks) == 0 {
		rinks = make([]*snc.Rink, 0)
	}
	c.JSON(http.StatusOK, rinks)
}

//func GetSpecificRink(c *gin.Context) {
//	teamID := req.PathParams["rinkID"]
//	if teamID == "" {
//		http.Error(w, "Missing rink ID", http.StatusBadRequest)
//	}
//	rID, err := strconv.ParseInt(teamID, 10, 32)
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	rink, err := ctx.RinkService.Rink(int(rID))
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//	encoder := json.NewEncoder(w)
//	err = encoder.Encode(rink)
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//}
//
//func UpdateRink(c *gin.Context) {
//	rinkID := req.PathParams["rinkID"]
//	if rinkID == "" {
//		http.Error(w, "Missing rink ID", http.StatusBadRequest)
//	}
//	rID, err := strconv.ParseInt(rinkID, 10, 32)
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	decoder := json.NewDecoder(req.Body)
//	defer req.Body.Close()
//	r := snc.Rink{}
//	err = decoder.Decode(&r)
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	if r.ID != 0 && r.ID != uint32(rID) {
//		msg := `There was a mismatch between ID specified in the path (URL) and the ID in the JSON provided.
//				They must be both, or you must omit the ID in the JSON and that in the path will be used`
//		log.Println(msg)
//		http.Error(w, msg, http.StatusBadRequest)
//		return
//	} else if r.ID == 0 {
//		// ID wasn't provided so assume the one in the URL
//		r.ID = uint32(rID)
//	}
//	// else case is the ID and path param ID match so proceed
//	err = ctx.RinkService.UpdateRink(&r)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//}
//
//------------------------------------------------------------------------------------------------//
// Divisions
//------------------------------------------------------------------------------------------------//
func CreateDivision(c *gin.Context) {
	s, err := services(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The server couldn't process the request"})
	}
	d := snc.Division{}
	if err = c.BindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if err = s.DivisionService.CreateDivision(&d); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

//func GetDivisions(c *gin.Context) {
//	divs, err := ctx.DivisionService.Divisions()
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//	encoder := json.NewEncoder(w)
//	encodeErr := encoder.Encode(divs)
//	if encodeErr != nil {
//		log.Println(encodeErr.Error())
//		http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
//		return
//	}
//}
//
//func GetSpecificDivision(c *gin.Context) {
//	divisionID := req.PathParams["divisionID"]
//	if divisionID == "" {
//		http.Error(w, "Missing division ID", http.StatusBadRequest)
//	}
//	dID, err := strconv.ParseInt(divisionID, 10, 32)
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	div, err := ctx.DivisionService.Division(int(dID))
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//	encoder := json.NewEncoder(w)
//	err = encoder.Encode(div)
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//}
//
//func UpdateDivision(c *gin.Context) {
//	divisionID := req.PathParams["divisionID"]
//	if divisionID == "" {
//		http.Error(w, "Missing division ID", http.StatusBadRequest)
//	}
//	dID, err := strconv.ParseInt(divisionID, 10, 32)
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	decoder := json.NewDecoder(req.Body)
//	defer req.Body.Close()
//	d := snc.Division{}
//	err = decoder.Decode(&d)
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	if d.ID != 0 && d.ID != uint32(dID) {
//		msg := `There was a mismatch between ID specified in the path (URL) and the ID in the JSON provided. They must be both, or you must omit the ID in the JSON and that in the path will be used`
//		log.Println(msg)
//		http.Error(w, msg, http.StatusBadRequest)
//		return
//	} else if d.ID == 0 {
//		// ID wasn't provided so assume the one in the URL
//		d.ID = uint32(dID)
//	}
//	// else case is the ID and path param ID match so proceed
//	err = ctx.DivisionService.UpdateDivision(&d)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//}
