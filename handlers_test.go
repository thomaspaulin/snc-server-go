package main

import (
	"os"
	"testing"
	"net/http/httptest"
	"github.com/thomaspaulin/snc-server-go/mock"
	"github.com/thomaspaulin/snc-server-go/snc"
	"net/http"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

// https://github.com/gin-gonic/gin/issues/55
// and
// https://github.com/gin-gonic/gin/pull/37/files
// for writing tests in Gin
func TestMain(m *testing.M) {
	// do the setup
	var retCode int
	// run the tests
	retCode	= m.Run()
	// do the tear down

	os.Exit(retCode)
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