package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/thomaspaulin/snc-server-go/mock"
	"github.com/thomaspaulin/snc-server-go/snc"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

// https://github.com/gin-gonic/gin/issues/55
// and
// https://github.com/gin-gonic/gin/pull/37/files
// for writing tests in Gin
func TestMain(m *testing.M) {
	// do the setup
	var retCode int
	// run the tests
	retCode = m.Run()
	// do the tear down

	os.Exit(retCode)
}

func TestGetMatches(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/matches", nil)

	matches := make([]*snc.Match, 2)
	matches[0] = &snc.Match{
		ID:        1,
		Start:     time.Date(2017, 4, 15, 4, 30, 0, 0, time.UTC),
		Division:  "C",
		Season:    2017,
		Status:    snc.Over,
		Away:      "Bears",
		Home:      "Tigers",
		AwayScore: 3,
		HomeScore: 1,
		Rink:      "Botany"}

	s := Services{}
	ms := mock.MatchService{}
	ms.MatchesFn = func() ([]*snc.Match, error) {
		return matches, nil
	}

	s.MatchService = &ms

	// this means the entire API has to be created again for each test :(
	r := APIEngine(s)
	r.ServeHTTP(w, req)

	assert.True(t, ms.MatchesInvoked)
	assert.Equal(t, 200, w.Code)

	actual := make([]*snc.Match, 1)
	json.NewDecoder(w.Body).Decode(&actual)

	// todo is this looking at memory address or contents?
	assert.Equal(t, matches, actual)
}

func TestGetTeams(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/teams", nil)

	teams := make([]*snc.Team, 2)
	teams[0] = &snc.Team{ID: 1, Division: "C", Name: "Bears"}
	teams[1] = &snc.Team{ID: 2, Division: "C", Name: "Hawks"}

	s := Services{}
	ts := mock.TeamService{}
	ts.TeamsFn = func() ([]*snc.Team, error) {
		return teams, nil
	}

	s.TeamService = &ts

	// this means the entire API has to be created again for each test :(
	r := APIEngine(s)
	r.ServeHTTP(w, req)

	assert.True(t, ts.TeamsInvoked)
	assert.Equal(t, 200, w.Code)

	actual := make([]*snc.Team, 2)
	json.NewDecoder(w.Body).Decode(&actual)

	// todo is this looking at memory address or contents?
	assert.Equal(t, teams, actual)
}

func TestGetSpecificRink(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/teams/1", nil)
	team := snc.Team{ID: 1, Division: "C", Name: "Bears"}

	s := Services{}
	ts := mock.TeamService{}
	ts.TeamFn = func(id int) (*snc.Team, error) {
		assert.Equal(t, id, 1)
		return &team, nil
	}

	s.TeamService = &ts

	// this means the entire API has to be created again for each test :(
	r := APIEngine(s)
	r.ServeHTTP(w, req)

	assert.True(t, ts.TeamInvoked)
	assert.Equal(t, 200, w.Code)

	actual := snc.Team{}
	json.NewDecoder(w.Body).Decode(&actual)
	assert.Equal(t, team, actual)
}

// todo uncomment when team creation is supported
//func TestCreateTeam(t *testing.T) {
//	team := snc.Team{Division: "C", Name: "Lions"}
//
//	w := httptest.NewRecorder()
//	j, _ := json.Marshal(team)
//	req := httptest.NewRequest("POST", "/v1/teams", bytes.NewBuffer(j))
//
//	s := Services{}
//	ts := mock.TeamService{}
//	ts.CreateFn = func(t *snc.Team) error {
//		return nil
//	}
//
//	s.TeamService = &ts
//	r := APIEngine(s)
//	r.ServeHTTP(w, req)
//
//	assert.True(t, ts.CreateInvoked)
//	assert.Equal(t, 200, w.Code)
//}
