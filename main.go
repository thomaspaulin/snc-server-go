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

var db *sql.DB

func connect(username string, password string , host string, dbName string) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, host, dbName)
	conn, _ := sql.Open("postgres", connStr)

	log.Printf("Here's what I'm using to connect to the database:\n" +
		"USER: %s\nHOST: %s\nDATABASE: %s", username, host, dbName)
	err := conn.Ping();
	if err != nil {
		panic(err)
	}
	db = conn
}

func (ctx *Context) getDBConnection(w web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	if ctx.database == nil {
		ctx.database = db
	}
	next(w, req)
}

func main() {
	connect(os.Getenv("SNC_USER"), os.Getenv("SNC_PW"), os.Getenv("SNC_HOST"), os.Getenv("SNC_DB"))
	r := web.New(Context{}).
			Middleware(web.LoggerMiddleware).
			Middleware(web.ShowErrorsMiddleware).
			Middleware((*Context).getDBConnection).
			Get("/", Index).
			Get("/hello", Hello).
			Get("/matches", (*Context).GetMatches).
			Post("/matches", (*Context).CreateMatches).
			Get("/matches/:matchID", (*Context).GetSpecificMatches).
			Get("/teams", (*Context).GetTeams).
			Get("/teams/:teamID", (*Context).GetSpecificTeam).
			Get("/rinks", (*Context).GetRinks).
			Get("/rinks/:rinkID", (*Context).GetSpecificRink)

	log.Printf("Starting up server on port %d\n", port())
	log.Fatal(http.ListenAndServe("localhost:4242", r))
}
