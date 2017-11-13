package snc

import (
	"time"
)

// todo handle the errors properly
const (
	PracticeGame = "PR"
	RegularGame  = "RS"
	PlayoffGame  = "PO"

	RegularGoal     = "RG"
	PowerPlayGoal   = "PP"
	ShortHandedGoal = "SH"

	Upcoming = "Upcoming"
	Underway = "Underway"
	Over     = "Over"
)

//-----------------------------------------------//
// Match
//-----------------------------------------------//
type Match struct {
	ID uint `json:"id"`
	// Datetime of the match start in UTC
	Start     time.Time `json:"start"`
	Season    int       `json:"season"`
	Status    string    `json:"status"`
	Division  Division  `json:"division"`
	Away      Team      `json:"away"`
	Home      Team      `json:"home"`
	AwayScore uint      `json:"awayScore"`
	HomeScore uint      `json:"homeScore"`
	Rink      Rink      `json:"rink"`
}

func CreateMatch(m Match) error {
	//id := 0
	//err := ms.DB.QueryRow(`
	//INSERT INTO matches
	// 	(start, season, status, away_id, home_id, away_score, home_score, rink_id)
	//VALUES
	//	($1, $2, $3, $4, $5, 0, 0, $6)
	//RETURNING match_id`, m.Start, m.Season, m.Status, awayID, homeID, m.AwayScore, m.HomeScore, rinkID).Scan(&id)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//return err
	return nil
}

//func FetchMatch(id int, DB *gorm.DB) (Match, error) {
//	m := Match{}
//	var divId, awayId, homeId, rinkId uint
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
//	FROM (matches AS m
//		INNER JOIN teams AS home
//			ON home.team_id = m.home_id
//		INNER JOIN teams AS away
//			ON away.team_id = m.away_id
//		INNER JOIN rinks AS rink
//			ON rink.rink_id = m.rink_id)
//	WHERE m.match_id = $1 AND m.deleted IS FALSE`, id).
//		Scan(&m.ID, &m.Start, &m.Season, divId, &m.Status, awayId, homeId, &m.AwayScore, &m.HomeScore, rinkId)
//	div, _ := FetchDivision(divId, DB)
//	m.Division = div
//	away, _ := FetchTeam(awayId, DB)
//	m.Away = away
//	home, _ := FetchTeam(homeId, DB)
//	m.Home = home
//	rink, _ := FetchRink(rinkId, DB)
//	m.Rink = rink
//	if err == sql.ErrNoRows {
//		return m, nil
//	} else if err != nil {
//		return Match{}, err
//	}
//	return m, nil
//}
//
//func FetchMatches(DB *gorm.DB) ([]Match, error) {
//	rows, err := DB.Query(`
//	SELECT
//		m.match_id,
//		m.start,
//		m.season,
//		m.division_id,
//		m.status,
//		m.away_id,
//		m.home_id,
//		m.away_score,
//		m.home_score,
//		m.rink_id
//	FROM matches AS m
//	WHERE m.deleted IS FALSE
//	ORDER BY m.start DESC`)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	matches := make([]Match, 0)
//	var divId, awayId, homeId, rinkId uint
//	for rows.Next() {
//		m := Match{}
//		err := rows.Scan(&m.ID, &m.Start, &m.Season, divId, &m.Status, awayId, homeId, &m.AwayScore, &m.HomeScore, rinkId)
//		// err here is the row.Scan(...) error. It shadows the err from outside the loop, and does not overwrite
//		if err != nil {
//			// probably the schema is wrong or the row is bad and so the database needs inspecting
//			// later on this might want to be changed to pass through and list the IDs of the bad rows
//			return nil, err
//		}
//		div, _ := FetchDivision(divId, DB)
//		m.Division = div
//		away, _ := FetchTeam(awayId, DB)
//		m.Away = away
//		home, _ := FetchTeam(homeId, DB)
//		m.Home = home
//		rink, _ := FetchRink(rinkId, DB)
//		m.Rink = rink
//		matches = append(matches, m)
//	}
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//	return matches, nil
//}
//
//func UpdateMatch(m Match, DB *gorm.DB) error {
//	// todo
//	return nil
//}
//
//func DeleteMatch(id int, DB *gorm.DB) error {
//	// todo
//	return nil
//}

//-----------------------------------------------//
// Goal
//-----------------------------------------------//
type Goal struct {
	ID       uint   `json:"id"`
	GoalType string `json:"goalType"`
	// ID of the team that scored
	Team   string `json:"team"`
	Period uint   `json:"period"`
	// Seconds left in the period when the goal was scored
	Time uint `json:"time"`
	// ID of the scoring player
	Scorer  string   `json:"scorer"`
	Assists []string `json:"assists"`
}

//-----------------------------------------------//
// Penalty
//-----------------------------------------------//
type Penalty struct {
	ID     uint   `json:"id"`
	Team   string `json:"team"`
	Period uint   `json:"period"`
	// Seconds left in the period when the penalty was incurred
	Time uint `json:"time"`
	// Name of the penalty
	Offense string `json:"offense"`
	// ID of the offender
	Offender string `json:"offender"`
	// Penalty Infraction Minutes
	PIM uint `json:"pim"`
}
