package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thomaspaulin/snc-server-go/snc"
	"log"
	"net/http"
	"strconv"
)

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
//func GetMatchesHandler(c *gin.Context) {
//	matches, err := snc.FetchMatches(DB)
//	if err != nil {
//		log.Println(err.Error())
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//	}
//
//	if len(matches) == 0 {
//		c.AbortWithStatus(http.StatusNoContent)
//	}
//	c.JSON(http.StatusOK, matches)
//}

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
func GetTeamsHandler(c *gin.Context) {
	teams, err := snc.FetchTeams(DB)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if len(teams) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.JSON(http.StatusOK, teams)
}

func GetSpecificTeamHandler(c *gin.Context) {
	var t snc.Team
	teamID := c.Param("teamID")
	if teamID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing team ID"})
	} else if _, err := strconv.Atoi(teamID); err != nil {
		t, err = snc.TeamCalled(teamID, DB)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		tID, _ := strconv.Atoi(teamID)
		t, err = snc.FetchTeam(uint(tID), DB)
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
func CreateRinkHandler(c *gin.Context) {
	r := snc.Rink{}
	if err := c.BindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if err = snc.CreateRink(r, DB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

func GetRinksHandler(c *gin.Context) {
	rinks, err := snc.FetchRinks(DB)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if len(rinks) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.JSON(http.StatusOK, rinks)
}

func GetSpecificRinkHandler(c *gin.Context) {
	rinkID := c.Param("rinkID")
	if rinkID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing rink ID"})
	} else if _, err := strconv.Atoi(rinkID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	rID, err := strconv.Atoi(rinkID)
	r, err := snc.FetchRink(uint(rID), DB)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, r)
}

func UpdateRinkHandler(c *gin.Context) {
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
		if err := snc.UpdateRink(r, DB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

//------------------------------------------------------------------------------------------------//
// Divisions
//------------------------------------------------------------------------------------------------//
func CreateDivisionHandler(c *gin.Context) {
	d := snc.Division{}
	if err := c.BindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if err = snc.CreateDivision(d, DB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

func GetDivisionsHandler(c *gin.Context) {
	divisions, err := snc.FetchDivisions(DB)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if len(divisions) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.JSON(http.StatusOK, divisions)
}

func GetSpecificDivisionHandler(c *gin.Context) {
	divisionID := c.Param("divisionID")
	if divisionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing division ID"})
	} else if _, err := strconv.Atoi(divisionID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	dID, err := strconv.Atoi(divisionID)
	d, err := snc.FetchDivision(uint(dID), DB)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, d)
}

func UpdateDivisionHandler(c *gin.Context) {
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
		if err := snc.UpdateDivision(d, DB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

//------------------------------------------------------------------------------------------------//
// Players
//------------------------------------------------------------------------------------------------//
func CreatePlayerHandler(c *gin.Context) {
	p := snc.Player{}
	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if err = snc.CreatePlayer(p, DB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

func GetPlayersHandler(c *gin.Context) {
	players, err := snc.FetchPlayers(DB)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if len(players) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.JSON(http.StatusOK, players)
}

func GetSpecificPlayerHandler(c *gin.Context) {
	playerID := c.Param("playerID")
	if playerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing division ID"})
	} else if _, err := strconv.Atoi(playerID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	pID, err := strconv.Atoi(playerID)
	d, err := snc.FetchPlayer(uint(pID), DB)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, d)
}

func UpdatePlayerHandler(c *gin.Context) {
	playerID := c.Param("playerID")
	if playerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing division ID"})
	} else if _, err := strconv.Atoi(playerID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	pID, err := strconv.Atoi(playerID)
	p := snc.Player{}
	if err = c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if p.ID != 0 && p.ID != uint(pID) {
			msg := `There was a mismatch between ID specified in the path (URL) and the ID in the JSON provided. They must be both, or you must omit the ID in the JSON and that in the path will be used`
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		} else if p.ID == 0 {
			// ID wasn't provided so assume the one in the URL
			p.ID = uint(pID)
		}
		if err := snc.UpdatePlayer(p, DB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}
