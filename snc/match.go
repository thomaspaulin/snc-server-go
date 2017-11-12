package snc

import (
	"time"
	"github.com/jinzhu/gorm"
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
	gorm.Model
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
	//Goals		[]*Goal		`json:"goals"`
	//Penalties	[]*Penalty	`json:"penalties"`
}

type MatchService interface {
	CreateMatch(m *Match) error
	Match(id int) (*Match, error)
	Matches() ([]*Match, error)
	UpdateMatch(m *Match) error
	DeleteMatch(id int) error
}

//-----------------------------------------------//
// Goal
//-----------------------------------------------//
type Goal struct {
	gorm.Model
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
	gorm.Model
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
