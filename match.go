package main

import (
	"time"
	"database/sql"
	"errors"
	"github.com/thomaspaulin/snc-server-go/database"
	"log"
)

// todo handle the errors properly
const (
	PracticeGame = 		"PR"
	RegularGame = 		"RS"
	PlayoffGame = 		"PO"

	RegularGoal = 		"RG"
	PowerPlayGoal = 	"PP"
	ShortHandedGoal = 	"SH"
)

// Models

type Match struct {
	id			uint32
	// Datetime of the match start in UTC
	start 		time.Time
	season		int
	away 		string
	home 		string
	awayScore	int
	homeScore	int
	rink		string
}

// TODO: unify the match and match summary classes? In a way they are the same thing
type MatchSummary struct {
	matchId		uint32
	// Datetime of the match start in UTC
	start		time.Time
	away		string
	home		string
	awayScore	int
	homeScore	int
	rink		string
	goals		[]*Goal
	penalties	[]*Penalty
	// todo:
	//  - shots (per team, per period)
	//  - power plays (per team, successful and total)
	//  - players and goalies indexed by team
}

type Goal struct {
	id			uint32
	goalType	string
	// ID of the team that scored
	team		string
	period		uint
	// Seconds left in the period when the goal was scored
	time		uint
	// ID of the scoring player
	scorer		string
	assists		[]string
}

type Penalty struct {
	id			uint32
	team		string
	period		uint
	// Seconds left in the period when the penalty was incurred
	time		uint
	// Name of the penalty
	offense		string
	// ID of the offender
	offender	string
	// Penalty Infraction Minutes
	pim			uint
}

// Database logic
func (m *Match) Save() (id uint32, err error) {
	id = 0
	if m.id > 0 {
		id, err = m.Create()
	} else {
		id, err = m.Update()
	}
	return id, err
}

func (m *Match) Create() (id uint32, err error) {
	database.DB.QueryRow("INSERT INTO matches " +
		"(start, season, away, home, awayScore, homeScore, rink) " +
		"VALUES " +
		"($1, $2, $3, $4, -1, -1, $5)" +
		"RETURNING match_id", m.start, m.season, m.away, m.home, m.awayScore, m.homeScore, m.rink).
		Scan(&id)
	if err != nil {
		log.Println(err.Error())
	}
	return id, nil
}

func FetchMatches() ([]*Match, error) {
	// TODO: Redo this query to do join and get the team names. It will fail as is
	rows, err := database.DB.Query("SELECT * FROM matches")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	matches := make([]*Match, 0)
	for rows.Next() {
		m := Match{}
		err := rows.Scan(&m.id, &m.start, &m.season, &m.away, &m.home, &m.awayScore, &m.homeScore, &m.rink)
		// err here is the row.Scan(...) error. It shadows the err from outside the loop, and does not overwrite
		if err != nil {
			// probably the schema is wrong or the row is bad and so the database needs inspecting
			// later on this might want to be changed to pass through and list the IDs of the bad rows
			return nil, err
		}
		matches = append(matches, &m)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return matches, nil
}

func FetchMatch(id uint32) (*Match, error) {
	m := Match{id: id}
	// TODO redo the query to do joins instead of selecting all columns
	err := database.DB.QueryRow("SELECT * WHERE id = $1", id).Scan(&m)
	if err == sql.ErrNoRows {
		return &Match{}, nil
	} else if err != nil {
		return nil, err
	}
	return &m, nil
}

func (m *Match) Update() (id uint32, err error) {
	return 0, errors.New("not implemented");
}
