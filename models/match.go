package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

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
	gorm.Model
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
