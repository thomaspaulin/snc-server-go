package models

import "time"

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