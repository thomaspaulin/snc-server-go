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
	"strings"
)

func port() int {
	p, present := os.LookupEnv("SNC_SERV_PORT")
	if !present {
		log.Println("Environment variable SNC_SERV_PORT not present, looking for PORT")
		p, present = os.LookupEnv("PORT")
		if !present {
			log.Println("Environment variable PORT not present, falling back to 4242")
			return 4242
		}
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		log.Println("Failed to parse the port to an integer. Using 4242 instead. The error was:")
		log.Println(err.Error())
		return 4242
	}
	return int(i)
}

func databaseURL() string {
	URL, present := os.LookupEnv("DATABASE_URL")
	if !present {
		log.Println("Environment variable 'DATABASE_URL' was not present, using the 'SNC_' prefixed environment variables to create the database URL.")
		username := os.Getenv("SNC_USER")
		password := os.Getenv("SNC_PW")
		host := os.Getenv("SNC_HOST")
		DBName := os.Getenv("SNC_DB")
		URL = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, host, DBName)
	}
	c := strings.LastIndex(URL, ":")
	a := strings.Index(URL, "@")
	u := URL[0:c+1] + "********" + URL[a:len(URL)]
	log.Printf("Here's what I'm using to connect to the database:\n%s", u)
	return URL
}

var DB *gorm.DB

func main() {
	var err error
	DB, err = gorm.Open("postgres", databaseURL())
	if err != nil {
		panic(err)
	}

	defer DB.Close()

	DB.AutoMigrate(&snc.Rink{}, &snc.Division{}, &snc.Team{}, &snc.Player{}, &snc.Match{}, &snc.Goal{}) //, &snc.Penalty{})
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
	port := port()
	log.Printf("main: starting up server on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
