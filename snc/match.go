package snc
//
//import (
//	"time"
//	"database/sql"
//	"errors"
//	"log"
//)
//
//// todo handle the errors properly
//const (
//	PracticeGame 	= 	"PR"
//	RegularGame 	= 	"RS"
//	PlayoffGame 	= 	"PO"
//
//	RegularGoal 	= 	"RG"
//	PowerPlayGoal 	= 	"PP"
//	ShortHandedGoal = 	"SH"
//
//	Upcoming 		=	"Upcoming"
//	Underway		=	"Underway"
//	Over			=	"Over"
//)
//
////-----------------------------------------------//
//// Match
////-----------------------------------------------//
//type Match struct {
//	ID			uint32		`json:"id"`
//	// Datetime of the match start in UTC
//	Start 		time.Time	`json:"start"`
//	Season		int			`json:"season"`
//	Status		string		`json:"status"`
//	Division	string		`json:"division"`
//	Away 		string		`json:"away"`
//	Home 		string		`json:"home"`
//	AwayScore	uint32		`json:"awayScore"`
//	HomeScore	uint32		`json:"homeScore"`
//	Rink		string		`json:"rink"`
//}
//// TODO UPDATE THE DATABASE SCHEMA SO THAT MATCHES HAVE A DIVISION_ID COLUMN
//
//
//func (m *Match) Create(DB *sql.DB) (id uint32, err error) {
//	// TODO: Save the teams and rinks first so we get the IDs and rink ID
//	awayID, err := teamID(m.Away, m.Division)
//	if err != nil {
//		log.Println("I was unable to find the away team so I tried saving but that failed too. I don't know how to handle this case.")
//		return 0, err
//	}
//	homeID, err := teamID(m.Home, m.Division)
//	if err != nil {
//		log.Println("I was unable to find the home team so I tried saving but that failed too. I don't know how to handle this case.")
//		return 0, err
//	}
//
//	rinkID, err := rinkID(m.Rink)
//	if err != nil {
//		return 0, err
//	}
//	// todo fix this
//
//	DB.QueryRow(`
//	INSERT INTO matches
//	 	(start, season, status, away_id, home_id, away_score, home_score, rink_id)
//	VALUES
//		($1, $2, $3, $4, $5, 0, 0, $6)
//	RETURNING match_id`, m.Start, m.Season, m.Status, awayID, homeID, m.AwayScore, m.HomeScore, rinkID).Scan(&id)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	return id, nil
//}
//
//func FetchMatches(DB *sql.DB) ([]*Match, error) {
//	rows, err := DB.Query(`
//	SELECT
//		m.match_id,
//		m.start,
//		m.season,
//		m.status,
//		away.name,
//		home.name,
//		m.away_score,
//		m.home_score,
//		rink.name
//	FROM (matches AS m INNER JOIN teams AS home ON home.team_id = m.home_id
//	INNER JOIN teams AS away ON away.team_id = m.away_id
//	INNER JOIN rinks AS rink ON rink.rink_id = m.rink_id)
//	ORDER BY m.start DESC`)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	matches := make([]*Match, 0)
//	for rows.Next() {
//		m := Match{}
//		err := rows.Scan(&m.ID, &m.Start, &m.Season, &m.Status, &m.Away, &m.Home, &m.AwayScore, &m.HomeScore, &m.Rink)
//		// err here is the row.Scan(...) error. It shadows the err from outside the loop, and does not overwrite
//		if err != nil {
//			// probably the schema is wrong or the row is bad and so the database needs inspecting
//			// later on this might want to be changed to pass through and list the IDs of the bad rows
//			return nil, err
//		}
//		matches = append(matches, &m)
//	}
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//	return matches, nil
//}
//
//func FetchMatch(DB *sql.DB, id uint32) (*Match, error) {
//	m := Match{ID: id}
//	err := DB.QueryRow(`
//	SELECT
//		m.match_id,
//		m.start,
//		m.season,
//		m.status,
//		away.name,
//		home.name,
//		m.away_score,
//		m.home_score,
//		rink.name
//	FROM (matches AS m INNER JOIN teams AS home ON home.team_id = m.home_id
//	INNER JOIN teams AS away ON away.team_id = m.away_id
//	INNER JOIN rinks AS rink ON rink.rink_id = m.rink_id)
//	WHERE match_id = $1`, id).Scan(&m.ID, &m.Start, &m.Season, &m.Status, &m.Away, &m.Home, &m.AwayScore, &m.HomeScore, &m.Rink)
//	if err == sql.ErrNoRows {
//		return &Match{}, nil
//	} else if err != nil {
//		return nil, err
//	}
//	return &m, nil
//}
//
//func (m *Match) Update(DB *sql.DB) (id uint32, err error) {
//	return 0, errors.New("not implemented");
//}
//
//// TODO: unify the match and match summary classes? In a way they are the same thing
////-----------------------------------------------//
//// Match Summary
////-----------------------------------------------//
//type MatchSummary struct {
//	MatchID		uint32		`json:"matchId"`
//	// Datetime of the match start in UTC
//	Start		time.Time	`json:"start"`
//	Away		string		`json:"away"`
//	Home		string		`json:"home"`
//	AwayScore	int			`json:"awayScore"`
//	HomeScore	int			`json:"homeScore"`
//	Rink		string		`json:"rink"`
//	Goals		[]*Goal		`json:"goals"`
//	Penalties	[]*Penalty	`json:"penalties"`
//	// todo:
//	//  - shots (per team, per period)
//	//  - power plays (per team, successful and total)
//	//  - players and goalies indexed by team
//}
//
////-----------------------------------------------//
//// Goal
////-----------------------------------------------//
//type Goal struct {
//	ID			uint32		`json:"id"`
//	GoalType	string		`json:"goalType"`
//	// ID of the team that scored
//	Team		string		`json:"team"`
//	Period		uint		`json:"period"`
//	// Seconds left in the period when the goal was scored
//	Time		uint		`json:"time"`
//	// ID of the scoring player
//	Scorer		string		`json:"scorer"`
//	Assists		[]string	`json:"assists"`
//}
//
////-----------------------------------------------//
//// Penalty
////-----------------------------------------------//
//type Penalty struct {
//	ID			uint32		`json:"id"`
//	Team		string		`json:"team"`
//	Period		uint		`json:"period"`
//	// Seconds left in the period when the penalty was incurred
//	Time		uint		`json:"time"`
//	// Name of the penalty
//	Offense		string		`json:"offense"`
//	// ID of the offender
//	Offender	string		`json:"offender"`
//	// Penalty Infraction Minutes
//	PIM			uint		`json:"pim"`
//}
