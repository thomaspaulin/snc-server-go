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

// Database logic
