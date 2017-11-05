package main

import (
	"log"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
	"strconv"
	"github.com/thomaspaulin/snc-server-go/postgres"
	"github.com/thomaspaulin/snc-server-go/snc"
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

type Context struct {
	DB 				*sql.DB

	// TODO these are temporary
	DivisionService snc.DivisionService
	RinkService		snc.RinkService
	TeamService 	snc.TeamService
}

var DB *sql.DB

func setDBConnection(ctx *gin.Context) {
	ctx.Set("DB", DB)
}

func initServices(DB *sql.DB) Services {
	s := Services{}

	ds := &postgres.DivisionService{DB: DB}
	s.DivisionService = ds

	ts := &postgres.TeamService{DB: DB}
	s.TeamService = ts

	rs := &postgres.RinkService{DB: DB}
	s.RinkService = rs

	return s
}

func APIEngine(s Services) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("services", &s)
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", Index)
		v1.GET("/hello", Hello)

		//Get("/matches", (*Context).GetMatches).
		//Post("/matches", (*Context).CreateMatches).
		//Get("/matches/:matchID", (*Context).GetSpecificMatches)
		v1.GET("/teams", GetTeams)
		v1.GET("/teams/:teamID", GetSpecificTeam)
		v1.GET("/rinks", GetRinks)
		v1.GET("/rinks/:rinkID", GetSpecificRink)
		v1.POST("/rinks", CreateRink)
		v1.PUT("/rinks/:rinkID", UpdateRink)
		v1.GET("/divisions", GetDivisions)
		v1.GET("/divisions/:divisionID", GetSpecificDivision)
		v1.POST("divisions", CreateDivision)
		v1.PUT("divisions/:divisionID", UpdateDivision)
	}
	return r
}

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

	router := APIEngine(initServices(DB))
	router.Use(setDBConnection)
	log.Printf("main: starting up server on port %d\n", port())
	log.Fatal(http.ListenAndServe("localhost:4242", router))
}