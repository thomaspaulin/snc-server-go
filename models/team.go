package models

import "github.com/jinzhu/gorm"

// Models

type Team struct {
	gorm.Model
	id			string
	name		string
	division	string
	logoURL		string
}

type Player struct {
	id			string
	number		uint
	name		string
	// players can be on multiple teams
	teams		[]string
	position	string
}

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

// Database logic
