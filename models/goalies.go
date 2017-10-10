package models

type Goalie struct {
	id			string
	number		uint
	name		string
	// goalies can be on multiple teams
	teams		[]string
	shots		uint
	saves		uint
	mins		uint
}