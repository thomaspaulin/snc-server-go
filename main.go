package main

import (
	"log"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
	"strconv"
	"github.com/thomaspaulin/snc-server-go/postgres"
	"github.com/thomaspaulin/snc-server-go/snc"
	"github.com/gocraft/web"
	"fmt"
	"net/http"
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

func (ctx *Context) getDBConnection(w web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	ctx.DB = DB
	next(w, req)
}

func (ctx *Context) initServices(w web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	ds := &postgres.DivisionService{DB: ctx.DB}
	ctx.DivisionService = ds

	ts := &postgres.TeamService{DB: ctx.DB}
	ctx.TeamService = ts

	rs := &postgres.RinkService{DB: ctx.DB}
	ctx.RinkService = rs

	next(w, req)
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

	r := web.New(Context{}).
			Middleware(web.LoggerMiddleware).
			Middleware(web.ShowErrorsMiddleware).
			Middleware((*Context).getDBConnection).
			Middleware((*Context).initServices).
			Get("/", Index).
			Get("/hello", Hello).
			//Get("/matches", (*Context).GetMatches).
			//Post("/matches", (*Context).CreateMatches).
			//Get("/matches/:matchID", (*Context).GetSpecificMatches).

			Get("/teams", (*Context).GetTeams).
			Get("/teams/:teamID", (*Context).GetSpecificTeam).

			Get("/rinks", (*Context).GetRinks).
			Get("/rinks/:rinkID", (*Context).GetSpecificRink).
			Post("/rinks", (*Context).CreateRink).
			Put("/rinks/:rinkID", (*Context).UpdateRink).

			Get("/divisions", (*Context).GetDivisions).
			Get("/divisions/:divisionID", (*Context).GetSpecificDivision).
			Post("divisions", (*Context).CreateDivision).
			Put("divisions/:divisionID", (*Context).UpdateDivision)

	log.Printf("main: starting up server on port %d\n", port())
	log.Fatal(http.ListenAndServe("localhost:4242", r))
}