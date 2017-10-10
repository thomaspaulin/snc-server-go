package models

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