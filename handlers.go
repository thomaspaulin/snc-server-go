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
	MatchService		snc.MatchService
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
func GetMatches(c *gin.Context) {
	s, err := services(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The server couldn't process the request"})
	}
	matches, err := s.MatchService.Matches()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if len(matches) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.JSON(http.StatusOK, matches)
}

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
//	match, err := FetchMatch(ctx.DB, uint(mID))
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
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.JSON(http.StatusOK, teams)
}

func GetSpecificTeam(c *gin.Context) {
	s, err := services(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The server couldn't process the request"})
	}
	var t *snc.Team
	teamID := c.Param("teamID")
	if teamID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing team ID"})
	} else if _, err := strconv.Atoi(teamID); err != nil {
		t, err = s.TeamService.TeamCalled(teamID)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		tID, _ := strconv.Atoi(teamID)
		t, err = s.TeamService.Team(int(tID))
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
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
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.JSON(http.StatusOK, rinks)
}

func GetSpecificRink(c *gin.Context) {
	s, err := services(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The server couldn't process the request"})
	}
	rinkID := c.Param("rinkID")
	if rinkID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing rink ID"})
	} else if _, err := strconv.Atoi(rinkID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	rID, err := strconv.Atoi(rinkID)
	r, err := s.RinkService.Rink(int(rID))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, r)
}

func UpdateRink(c *gin.Context) {
	s, err := services(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The server couldn't process the request"})
	}
	rinkID := c.Param("rinkID")
	if rinkID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing rink ID"})
	} else if _, err := strconv.Atoi(rinkID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	rID, err := strconv.Atoi(rinkID)
	r := snc.Rink{}
	if err = c.BindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if r.ID != 0 && r.ID != uint(rID) {
			msg := `There was a mismatch between ID specified in the path (URL) and the ID in the JSON provided. They must be both, or you must omit the ID in the JSON and that in the path will be used`
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		} else if r.ID == 0 {
			// ID wasn't provided so assume the one in the URL
			r.ID = uint(rID)
		}
		if err := s.RinkService.UpdateRink(&r); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

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

func GetDivisions(c *gin.Context) {
	s, err := services(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The server couldn't process the request"})
	}
	divisions, err := s.DivisionService.Divisions()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if len(divisions) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.JSON(http.StatusOK, divisions)
}

func GetSpecificDivision(c *gin.Context) {
	s, err := services(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The server couldn't process the request"})
	}
	divisionID := c.Param("divisionID")
	if divisionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing division ID"})
	} else if _, err := strconv.Atoi(divisionID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	dID, err := strconv.Atoi(divisionID)
	d, err := s.DivisionService.Division(int(dID))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, d)
}

func UpdateDivision(c *gin.Context) {
	s, err := services(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The server couldn't process the request"})
	}
	divisionID := c.Param("divisionID")
	if divisionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing division ID"})
	} else if _, err := strconv.Atoi(divisionID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	dID, err := strconv.Atoi(divisionID)
	d := snc.Division{}
	if err = c.BindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if d.ID != 0 && d.ID != uint(dID) {
			msg := `There was a mismatch between ID specified in the path (URL) and the ID in the JSON provided. They must be both, or you must omit the ID in the JSON and that in the path will be used`
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		} else if d.ID == 0 {
			// ID wasn't provided so assume the one in the URL
			d.ID = uint(dID)
		}
		if err := s.DivisionService.UpdateDivision(&d); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}
