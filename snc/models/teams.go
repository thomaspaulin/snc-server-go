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

// Database logic
