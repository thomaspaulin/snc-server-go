package snc

import (
	"errors"
	"github.com/jinzhu/gorm"
)

//--------------------------------------------------------------------------------------------------------------------//
// Team
//--------------------------------------------------------------------------------------------------------------------//
type Team struct {
	gorm.Model
	Name		string	`json:"name"`
	Division	string	`json:"division"`
	LogoURL		string	`json:"logoURL"`
}

type TeamService interface {
	CreateTeam(t *Team) error
	Team(id int) (*Team, error)
	Teams() ([]*Team, error)
	UpdateTeam(t *Team) error
	DeleteTeam(id int) error
}

var ErrMultipleTeams = errors.New("snc: expected only one team but got multiple")

//--------------------------------------------------------------------------------------------------------------------//
// Division
//--------------------------------------------------------------------------------------------------------------------//
type Division struct {
	gorm.Model
	Name		string	`json:"name"`
}

type DivisionService interface {
	CreateDivision(d *Division) error
	Division(id int) (*Division, error)
	Divisions() ([]*Division, error)
	UpdateDivision(d *Division) error
	DeleteDivision(id int) error
}

//--------------------------------------------------------------------------------------------------------------------//
// Player
//--------------------------------------------------------------------------------------------------------------------//
type Player struct {
	gorm.Model
	Number		uint		`json:"number"`
	Name		string		`json:"name"`
	Teams		[]string	`json:"teams"`
	Position	string		`json:"position"`
}

//--------------------------------------------------------------------------------------------------------------------//
// Goalie
//--------------------------------------------------------------------------------------------------------------------//
type Goalie struct {
	gorm.Model
	Number		uint		`json:"number"`
	Name		string		`json:"name"`
	Teams		[]string	`json:"teams"`
	Shots		uint		`json:"shots"`
	Saves		uint		`json:"saves"`
	Mins		uint		`json:"mins"`
}
