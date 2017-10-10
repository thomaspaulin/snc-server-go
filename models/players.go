package models

// Models

type Player struct {
	id			string
	number		uint
	name		string
	// players can be on multiple teams
	teams		[]string
	position	string
}

// Database Logic
