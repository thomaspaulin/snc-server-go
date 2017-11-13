package main

import (
	"log"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
	"strconv"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func port() int {
	p := os.Getenv("SNC_SERV_PORT")
	i, err := strconv.ParseInt(p, 10, 32)
	if err == nil {
		return int(i)
	}
	return 4242
}

var DB *sql.DB

func main() {
	username := os.Getenv("SNC_USER")
	password := os.Getenv("SNC_PW")
	host := os.Getenv("SNC_HOST")
	DBName := os.Getenv("SNC_DB")

	var err error
	DB, err = sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, host, DBName))
	if err != nil {
		panic(err)
	}
	log.Printf("Here's what I'm using to connect to the database:\n" +
		"USER: %s\nHOST: %s\nDATABASE: %s", username, host, DBName)
	err = DB.Ping();
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	// todo find a way to do it such that the services aren't in the context
	// todo rename updateX to be more in line with replaceX
	// todo use PATCH to update parts of an entity and

	r := gin.Default()

	v1 := r.Group("/api/v0")
	{
		v1.GET("/", Index)
		v1.GET("/hello", Hello)

		v1.GET("/matches", GetMatchesHandler)
		//Post("/matches", (*Context).CreateMatches).
		//Get("/matches/:matchID", (*Context).GetSpecificMatches)
		v1.GET("/teams", GetTeamsHandler)
		v1.GET("/teams/:teamID", GetSpecificTeamHandler)
		v1.GET("/rinks", GetRinksHandler)
		v1.GET("/rinks/:rinkID", GetSpecificRinkHandler)
		v1.POST("/rinks", CreateRinkHandler)
		v1.PUT("/rinks/:rinkID", UpdateRinkHandler)
		v1.GET("/divisions", GetDivisionsHandler)
		v1.GET("/divisions/:divisionID", GetSpecificDivisionHandler)
		v1.POST("divisions", CreateDivisionHandler)
		v1.PUT("divisions/:divisionID", UpdateDivisionHandler)
	}
	log.Printf("main: starting up server on port %d\n", port())
	log.Fatal(http.ListenAndServe("localhost:4242", r))
}