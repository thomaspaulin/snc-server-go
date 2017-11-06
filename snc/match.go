package snc

import (
	"time"
)

// todo handle the errors properly
const (
	PracticeGame 	= 	"PR"
	RegularGame 	= 	"RS"
	PlayoffGame 	= 	"PO"

	RegularGoal 	= 	"RG"
	PowerPlayGoal 	= 	"PP"
	ShortHandedGoal = 	"SH"

	Upcoming 		=	"Upcoming"
	Underway		=	"Underway"
	Over			=	"Over"
)

//-----------------------------------------------//
// Match
//-----------------------------------------------//
type Match struct {
	ID			uint32		`json:"id"`
	// Datetime of the match start in UTC
	Start 		time.Time	`json:"start"`
	Season		int			`json:"season"`
	Status		string		`json:"status"`
	Division	string		`json:"division"`
	Away 		string		`json:"away"`
	Home 		string		`json:"home"`
	AwayScore	uint32		`json:"awayScore"`
	HomeScore	uint32		`json:"homeScore"`
	Rink		string		`json:"rink"`
}

type MatchService interface {
	CreateMatch(m *Match) error
	Match(id int) (*Match, error)
	Matches() ([]*Match, error)
	UpdateMatch(m *Match) error
	DeleteMatch(id int) error
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
