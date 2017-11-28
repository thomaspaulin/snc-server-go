package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thomaspaulin/snc-server-go/snc"
	"log"
	"net/http"
	"strconv"
	"fmt"
)

func Index(c *gin.Context) {
	c.Redirect(307, "/api/v0/")
}

func APIIndex(c *gin.Context) {
	c.String(200, "SNC API is a work in progress.")
}

//------------------------------------------------------------------------------------------------//
// Matches
//------------------------------------------------------------------------------------------------//
func CreateMatchHandler(c *gin.Context) {
	m := snc.Match{}
	if err := c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		log.Printf("%v\n", m)
		if err = snc.CreateMatch(m, DB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

func GetMatchesHandler(c *gin.Context) {
	var pagination snc.Pagination
	if c.ShouldBindQuery(&pagination) != nil {
		fmt.Println("Failed to bind query params")
		pagination.Limit = -1
		pagination.Offset = -1
	}
	matches, err := snc.FetchMatches(pagination, DB)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if len(matches) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.JSON(http.StatusOK, matches)
}

func GetSpecificMatchHandler(c *gin.Context) {
	matchID := c.Param("matchID")
	if matchID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing rink ID"})
	} else if _, err := strconv.Atoi(matchID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	mID, err := strconv.Atoi(matchID)
	m, err := snc.FetchMatch(uint(mID), DB)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if m.ID != 0 {
		c.JSON(http.StatusOK, m)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func UpdateMatchHandler(c *gin.Context) {
	matchID := c.Param("matchID")
	if matchID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing match ID"})
	} else if _, err := strconv.Atoi(matchID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	tID, err := strconv.Atoi(matchID)
	m := snc.Match{}
	if err = c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if m.ID != 0 && m.ID != uint(tID) {
			msg := `There was a mismatch between ID specified in the path (URL) and the ID in the JSON provided. They must be both, or you must omit the ID in the JSON and that in the path will be used`
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		} else if m.ID == 0 {
			// ID wasn't provided so assume the one in the URL
			m.ID = uint(tID)
		}
		if err := snc.UpdateMatch(m, DB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

func DeleteMatchHandler(c *gin.Context) {
	matchID := c.Param("matchID")
	if matchID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing match ID"})
	} else if _, err := strconv.Atoi(matchID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	mID, err := strconv.Atoi(matchID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		if err := snc.DeleteMatch(uint(mID), DB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

//------------------------------------------------------------------------------------------------//
// Goals
//------------------------------------------------------------------------------------------------//
func CreateGoalHandler(c *gin.Context) {
	g := snc.Goal{}
	if err := c.BindJSON(&g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		log.Printf("%v\n", g)
		if err = snc.CreateGoal(g, DB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

func GetGoalsHandler(c *gin.Context) {
	goals, err := snc.FetchGoals(DB)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if len(goals) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.JSON(http.StatusOK, goals)
}

func GetSpecificGoalHandler(c *gin.Context) {
	goalID := c.Param("goalID")
	if goalID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing rink ID"})
	} else if _, err := strconv.Atoi(goalID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	gID, err := strconv.Atoi(goalID)
	g, err := snc.FetchGoal(uint(gID), DB)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if g.ID != 0 {
		c.JSON(http.StatusOK, g)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func UpdateGoalHandler(c *gin.Context) {
	goalID := c.Param("goalID")
	if goalID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing goal ID"})
	} else if _, err := strconv.Atoi(goalID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	tID, err := strconv.Atoi(goalID)
	g := snc.Goal{}
	if err = c.BindJSON(&g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if g.ID != 0 && g.ID != uint(tID) {
			msg := `There was a mismatch between ID specified in the path (URL) and the ID in the JSON provided. They must be both, or you must omit the ID in the JSON and that in the path will be used`
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		} else if g.ID == 0 {
			// ID wasn't provided so assume the one in the URL
			g.ID = uint(tID)
		}
		if err := snc.UpdateGoal(g, DB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

func DeleteGoalHandler(c *gin.Context) {
	goalID := c.Param("goalID")
	if goalID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing goal ID"})
	} else if _, err := strconv.Atoi(goalID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	gID, err := strconv.Atoi(goalID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		if err := snc.DeleteGoal(uint(gID), DB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

//------------------------------------------------------------------------------------------------//
// Teams
//------------------------------------------------------------------------------------------------//
func CreateTeamHandler(c *gin.Context) {
	t := snc.Team{}
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if err = snc.CreateTeam(t, DB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

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
		t, err = snc.FetchTeamNamed(teamID, DB)
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
	if t.ID != 0 {
		c.JSON(http.StatusOK, t)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func UpdateTeamHandler(c *gin.Context) {
	teamID := c.Param("teamID")
	if teamID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing team ID"})
	} else if _, err := strconv.Atoi(teamID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	tID, err := strconv.Atoi(teamID)
	t := snc.Team{}
	if err = c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if t.ID != 0 && t.ID != uint(tID) {
			msg := `There was a mismatch between ID specified in the path (URL) and the ID in the JSON provided. They must be both, or you must omit the ID in the JSON and that in the path will be used`
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		} else if t.ID == 0 {
			// ID wasn't provided so assume the one in the URL
			t.ID = uint(tID)
		}
		if err := snc.UpdateTeam(t, DB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

func DeleteTeamHandler(c *gin.Context) {
	teamID := c.Param("teamID")
	if teamID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing team ID"})
	} else if _, err := strconv.Atoi(teamID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	tID, err := strconv.Atoi(teamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		if err := snc.DeleteTeam(uint(tID), DB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
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
	if r.ID != 0 {
		c.JSON(http.StatusOK, r)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
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

func DeleteRinkHandler(c *gin.Context) {
	rinkID := c.Param("rinkID")
	if rinkID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing rink ID"})
	} else if _, err := strconv.Atoi(rinkID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	rID, err := strconv.Atoi(rinkID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if err := snc.DeleteRink(rID, DB); err != nil {
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
	if d.ID != 0 {
		c.JSON(http.StatusOK, d)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
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

func DeleteDivisionHandler(c *gin.Context) {
	divisionID := c.Param("divisionID")
	if divisionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing division ID"})
	} else if _, err := strconv.Atoi(divisionID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	dID, err := strconv.Atoi(divisionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if err := snc.DeleteDivision(dID, DB); err != nil {
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
	p, err := snc.FetchPlayer(uint(pID), DB)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if p.ID != 0 {
		c.JSON(http.StatusOK, p)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
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

func DeletePlayerHandler(c *gin.Context) {
	playerID := c.Param("playerID")
	if playerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing division ID"})
	} else if _, err := strconv.Atoi(playerID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer greater than 0"})
	}
	pID, err := strconv.Atoi(playerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		if err := snc.DeletePlayer(pID, DB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}
