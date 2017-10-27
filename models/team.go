package models

// Models

type Team struct {
	id			uint32
	name		string
	division	string
	logoURL		string
}

type Division struct {
	id			uint32
	name		string
}

type Player struct {
	id			uint32
	number		uint
	name		string
	// players can be on multiple teams
	teams		[]string
	position	string
}

type Goalie struct {
	id			uint32
	number		uint
	name		string
	teams		[]string
	shots		uint
	saves		uint
	mins		uint
}

// Database logic
