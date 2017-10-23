package models

import (
	"time"
	"database/sql"
	"errors"
)

// todo handle the errors properly
const (
	PracticeGame = "PR"
	RegularGame = "RS"
	PlayoffGame = "PO"

	RegularGoal = "RG"
	PowerPlayGoal = "PP"
	ShortHandedGoal = "SH"
)

// Models

type Match struct {
	id			uint32
	// Datetime of the match start in UTC
	start 		time.Time
	season		int
	// ID of the away team
	away 		string
	// ID of the home team
	home 		string
	awayScore	int
	homeScore	int
	rink		string
}

type MatchSummary struct {
	// Datetime of the match start in UTC
	start		time.Time
	// ID of the away team
	away		string
	// ID of the home team
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
	goalType	string
	// ID of the team that scored
	team		string
	period		uint
	// Seconds left in the period when the goal was scored
	time		uint
	// ID of the scoring player
	scorer		string
	// IDs of the players that assisted
	assists		[]string
}

type Penalty struct {
	// ID of the team that was penalised
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
func (m *Match) Save(db *sql.DB) error {
	err := m.create(db)
	if err != nil {
		return m.update(db)
	} else {
		return nil
	}
}

func FetchMatches(db *sql.DB) ([]*Match, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	rows, err := tx.Query("SELECT * FROM matches")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	matches := make([]*Match, 0)
	for rows.Next() {
		m := new(Match)
		err := rows.Scan(&m.id, &m.start, &m.season, &m.away, &m.home, &m.awayScore, &m.homeScore, &m.rink)
		// err here is the row.Scan(...) error. It shadows the err from outside the loop, and does not overwrite
		if err != nil {
			// probably the schema is wrong or the row is bad and so the database needs inspecting
			// later on this might want to be changed to pass through and list the IDs of the bad rows
			return nil, err
		}
		matches = append(matches, m)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func (m *Match) create(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO matches (start, season, away, home, awayScore, homeScore, rink) VALUES (?, ?, ?, ?, -1, -1, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(m.start, m.season, m.awayScore, m.home, m.rink)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (m *Match) update(db *sql.DB) error {
	return errors.New("not implemented");
}