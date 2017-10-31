package main

import (
	"time"
	"database/sql"
	"errors"
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

//-----------------------------------------------//
// Match
//-----------------------------------------------//
type Match struct {
	ID			uint32		`json:"id"`
	// Datetime of the match start in UTC
	Start 		time.Time	`json:"start"`
	Season		int			`json:"season"`
	Away 		string		`json:"away"`
	Home 		string		`json:"home"`
	AwayScore	int			`json:"awayScore"`
	HomeScore	int			`json:"homeScore"`
	Rink		string		`json:"rink"`
}

func (m *Match) Save(db *sql.DB) (id uint32, err error) {
	id = 0
	if m.ID > 0 {
		id, err = m.Create(db)
	} else {
		id, err = m.Update(db)
	}
	return id, err
}

func (m *Match) Create(db *sql.DB) (id uint32, err error) {
	db.QueryRow("INSERT INTO matches " +
		"(start, season, away, home, awayScore, homeScore, rink) " +
		"VALUES " +
		"($1, $2, $3, $4, -1, -1, $5)" +
		"RETURNING match_id", m.Start, m.Season, m.Away, m.Home, m.AwayScore, m.HomeScore, m.Rink).
		Scan(&id)
	if err != nil {
		log.Println(err.Error())
	}
	return id, nil
}

func FetchMatches(db *sql.DB) ([]*Match, error) {
	// TODO: Redo this query to do join and get the team names. It will fail as is
	rows, err := db.Query("SELECT * FROM matches")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	matches := make([]*Match, 0)
	for rows.Next() {
		m := Match{}
		err := rows.Scan(&m.ID, &m.Start, &m.Season, &m.Away, &m.Home, &m.AwayScore, &m.HomeScore, &m.Rink)
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

func FetchMatch(db *sql.DB, id uint32) (*Match, error) {
	m := Match{ID: id}
	// TODO redo the query to do joins instead of selecting all columns
	err := db.QueryRow("SELECT * WHERE id = $1", id).Scan(&m)
	if err == sql.ErrNoRows {
		return &Match{}, nil
	} else if err != nil {
		return nil, err
	}
	return &m, nil
}

func (m *Match) Update(db *sql.DB) (id uint32, err error) {
	return 0, errors.New("not implemented");
}

// TODO: unify the match and match summary classes? In a way they are the same thing
//-----------------------------------------------//
// Match Summary
//-----------------------------------------------//
type MatchSummary struct {
	MatchID		uint32		`json:"matchId"`
	// Datetime of the match start in UTC
	Start		time.Time	`json:"start"`
	Away		string		`json:"away"`
	Home		string		`json:"home"`
	AwayScore	int			`json:"awayScore"`
	HomeScore	int			`json:"homeScore"`
	Rink		string		`json:"rink"`
	Goals		[]*Goal		`json:"goals"`
	Penalties	[]*Penalty	`json:"penalties"`
	// todo:
	//  - shots (per team, per period)
	//  - power plays (per team, successful and total)
	//  - players and goalies indexed by team
}

//-----------------------------------------------//
// Goal
//-----------------------------------------------//
type Goal struct {
	ID			uint32		`json:"id"`
	GoalType	string		`json:"goalType"`
	// ID of the team that scored
	Team		string		`json:"team"`
	Period		uint		`json:"period"`
	// Seconds left in the period when the goal was scored
	Time		uint		`json:"time"`
	// ID of the scoring player
	Scorer		string		`json:"scorer"`
	Assists		[]string	`json:"assists"`
}

//-----------------------------------------------//
// Penalty
//-----------------------------------------------//
type Penalty struct {
	ID			uint32		`json:"id"`
	Team		string		`json:"team"`
	Period		uint		`json:"period"`
	// Seconds left in the period when the penalty was incurred
	Time		uint		`json:"time"`
	// Name of the penalty
	Offense		string		`json:"offense"`
	// ID of the offender
	Offender	string		`json:"offender"`
	// Penalty Infraction Minutes
	PIM			uint		`json:"pim"`
}
