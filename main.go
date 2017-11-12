package main

import (
	"log"
	"os"
	_ "github.com/lib/pq"
	"strconv"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/thomaspaulin/snc-server-go/snc"
)

func port() int {
	p := os.Getenv("SNC_SERV_PORT")
	i, err := strconv.ParseInt(p, 10, 32)
	if err == nil {
		return int(i)
	}
	return 4242
}

//var DB *sql.DB
var DB *gorm.DB

func APIEngine() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{

		//v1.GET("/matches", GetMatches)
		////Post("/matches", (*Context).CreateMatches).
		////Get("/matches/:matchID", (*Context).GetSpecificMatches)
		//v1.GET("/teams", GetTeams)
		//v1.GET("/teams/:teamID", GetSpecificTeam)
		//v1.GET("/rinks", GetRinks)
		//v1.GET("/rinks/:rinkID", GetSpecificRink)
		//v1.POST("/rinks", CreateRink)
		//v1.PUT("/rinks/:rinkID", UpdateRink)
		v1.GET("/divisions", GetDivisions)
		//v1.GET("/divisions/:divisionID", GetSpecificDivision)
		v1.POST("divisions", CreateDivision)
		//v1.PUT("divisions/:divisionID", UpdateDivision)
	}
	return r
}

func main() {
	username := os.Getenv("SNC_USER")
	password := os.Getenv("SNC_PW")
	host := os.Getenv("SNC_HOST")
	DBName := "snc_gorm"//os.Getenv("SNC_DB")

	var err error
	//DB, err = sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, host, DBName))
	DB, err = gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, host, DBName))
	if err != nil {
		panic(err)
	}
	log.Printf("Here's what I'm using to connect to the database:\n" +
		"USER: %s\nHOST: %s\nDATABASE: %s", username, host, DBName)
	//err = DB.Ping();
	//if err != nil {
	//	panic(err)
	//}
	defer DB.Close()
	DB.AutoMigrate(&snc.Division{})

	// todo find a way to do it such that the services aren't in the context
	// todo rename updateX to be more in line with replaceX
	// todo use PATCH to update parts of an entity and

	router := APIEngine()
	log.Printf("main: starting up server on port %d\n", port())
	log.Fatal(http.ListenAndServe("localhost:4242", router))
}