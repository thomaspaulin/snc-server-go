package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/gocraft/web"
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

type Context struct {
	database *sql.DB
}

func (ctx *Context) ConnectToDB(w web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	username := os.Getenv("SNC_USER")
	host := os.Getenv("SNC_HOST")
	DBName := os.Getenv("SNC_DB")

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s", username, os.Getenv("SNC_PW"), host, DBName)
	db, _ := sql.Open("postgres", connStr)

	log.Printf("Here's what I'm using to connect to the database:\n" +
		"USER: %s\nHOST: %s\nDATABASE: %s", username, host, DBName)
	err := db.Ping();
	if err != nil {
		panic(err)
	}
	ctx.database = db
	next(w, req)
}

func main() {
	r := web.New(Context{}).
			Middleware(web.LoggerMiddleware).
			Middleware(web.ShowErrorsMiddleware).
			Middleware((*Context).ConnectToDB).
			Get("/", Index).
			Get("/hello", Hello).
			Get("/matches", (*Context).GetMatches).
			Post("/matches", (*Context).CreateMatches).
			Get("/matches/:matchID", (*Context).GetSpecificMatches).
			Get("/teams/:teamID", (*Context).GetSpecificTeam)

	log.Printf("Starting up server on port %d\n", port())
	log.Fatal(http.ListenAndServe("localhost:4242", r))
}
