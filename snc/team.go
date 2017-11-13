package snc

import (
	"errors"
)

//--------------------------------------------------------------------------------------------------------------------//
// Team
//--------------------------------------------------------------------------------------------------------------------//
type Team struct {
	ID			uint	`json:"id"`
	Name		string	`json:"name"`
	Division	string	`json:"division"`
	LogoURL		string	`json:"logoURL"`
}

// todo use copy rather than pointer
type TeamService interface {
	CreateTeam(t *Team) error
	Team(id int) (*Team, error)
	TeamCalled(name string) (*Team, error)
	Teams() ([]*Team, error)
	UpdateTeam(t *Team) error
	DeleteTeam(id int) error
}

var ErrMultipleTeams = errors.New("snc: expected only one team but got multiple")

//--------------------------------------------------------------------------------------------------------------------//
// Division
//--------------------------------------------------------------------------------------------------------------------//
type Division struct {
	ID			uint	`json:"id"`
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
// Player (Goalies included)
//--------------------------------------------------------------------------------------------------------------------//
type Player struct {
	ID			uint		`json:"id"`
	Number		uint		`json:"number"`
	Name		string		`json:"name"`
	Teams		[]string	`json:"teams"`
	Position	string		`json:"position"`
}
