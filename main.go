package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/thomaspaulin/snc-server-go/snc"
	"log"
	"net/http"
	"os"
	"strconv"
)

func port() int {
	p := os.Getenv("SNC_SERV_PORT")
	i, err := strconv.ParseInt(p, 10, 32)
	if err == nil {
		return int(i)
	}
	return 4242
}

var DB *gorm.DB

func main() {
	username := os.Getenv("SNC_USER")
	password := os.Getenv("SNC_PW")
	host := os.Getenv("SNC_HOST")
	DBName := "snc_gorm" //os.Getenv("SNC_DB")

	var err error
	DB, err = gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, host, DBName))
	if err != nil {
		panic(err)
	}
	log.Printf("Here's what I'm using to connect to the database:\n"+
		"USER: %s\nHOST: %s\nDATABASE: %s", username, host, DBName)
	defer DB.Close()

	DB.AutoMigrate(&snc.Rink{}, &snc.Division{}, &snc.Team{}, &snc.Player{}, &snc.Match{}, &snc.Goal{})//, &snc.Penalty{})
	// todo find a way to do it such that the services aren't in the context
	// todo rename updateX to be more in line with replaceX
	// todo use PATCH to update parts of an entity and
	r := gin.Default()

	v1 := r.Group("/api/v0")
	{
		v1.GET("/", Index)
		v1.GET("/hello", Hello)

		v1.GET("/matches", GetMatchesHandler)
		v1.GET("/matches/:matchID", GetSpecificMatchHandler)
		v1.POST("/matches", CreateMatchHandler)
		v1.PUT("/matches/:matchID", UpdateMatchHandler)
		v1.DELETE("/matches/:matchID", DeleteMatchHandler)

		v1.GET("/goals", GetGoalsHandler)
		v1.GET("/goals/:goalID", GetSpecificGoalHandler)
		v1.POST("/goals", CreateGoalHandler)
		v1.PUT("/goals/:goalID", UpdateGoalHandler)
		v1.DELETE("/goals/:goalID", DeleteGoalHandler)

		v1.GET("/teams", GetTeamsHandler)
		v1.GET("/teams/:teamID", GetSpecificTeamHandler)
		v1.POST("/teams", CreateTeamHandler)
		v1.PUT("/teams/:teamID", UpdateTeamHandler)
		v1.DELETE("/teams/:teamID", DeleteTeamHandler)

		v1.GET("/rinks", GetRinksHandler)
		v1.GET("/rinks/:rinkID", GetSpecificRinkHandler)
		v1.POST("/rinks", CreateRinkHandler)
		v1.PUT("/rinks/:rinkID", UpdateRinkHandler)
		v1.DELETE("/rinks/:rinkID", DeleteDivisionHandler)

		v1.GET("/divisions", GetDivisionsHandler)
		v1.GET("/divisions/:divisionID", GetSpecificDivisionHandler)
		v1.POST("/divisions", CreateDivisionHandler)
		v1.PUT("/divisions/:divisionID", UpdateDivisionHandler)
		v1.DELETE("/divisions/:divisionID", DeleteDivisionHandler)

		v1.GET("/players", GetPlayersHandler)
		v1.GET("/players/:playerID", GetSpecificPlayerHandler)
		v1.POST("/players", CreatePlayerHandler)
		v1.PUT("/players/:playerID", UpdatePlayerHandler)
		v1.DELETE("/players/:playerID", DeletePlayerHandler)
	}
	log.Printf("main: starting up server on port %d\n", port())
	log.Fatal(http.ListenAndServe("localhost:4242", r))
}
