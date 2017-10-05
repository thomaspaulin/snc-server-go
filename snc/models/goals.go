package models

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
