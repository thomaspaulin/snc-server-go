package main

import (
	"github.com/thomaspaulin/snc-server-go/database"
	"database/sql"
	"fmt"
	"log"
)

//-----------------------------------------------//
// Team
//-----------------------------------------------//
type Team struct {
	id			uint32	`json:"id"`
	name		string	`json:"name"`
	division	string	`json:"division"`
	logoURL		string	`json:"logoURL"`
}

func (t *Team) Create() (id uint32, err error) {
	d := Division{name: t.division}
	divID, err := d.Save()
	id = 0
	err = database.DB.QueryRow("INSERT INTO teams " +
									"(name, division_id, logo_url) " +
								"VALUES " +
									"(?, ?, ?) " +
								"RETURNING team_id", t.name, divID, t.logoURL).Scan(&id)
	if err != nil {
		log.Println(err.Error())
	}
	return id, nil
}

func FetchTeamByID(id uint32) (*Team, error) {
	t := Team{id: id}
	err := database.DB.QueryRow("SELECT teams.name AS team_name, " +
								"divisions.name AS div_name, " +
								"teams.logo_url " +
							"FROM teams " +
							"JOIN divisions " +
								"ON teams.division_id = divisions.division_id " +
							"WHERE team_id = ?", id).Scan(&t.id, &t.division, &t.logoURL)
	if err == sql.ErrNoRows {
		return &Team{}, nil
	} else if err != nil {
		return nil, err
	}
	return &t, nil
}

func FetchTeam(teamName string, divName string) (*Team, error) {
	rows, err := database.DB.Query("SELECT teams.team_id AS team_id, " +
									"teams.name AS team_name, " +
									"divisions.name AS div_name, " +
									"teams.logo_url " +
								"FROM teams " +
								"JOIN divisions " +
									"ON teams.division_id = divisions.division_id " +
								"WHERE teams.name = ? " +
									"AND divisions.name = ?", teamName, divName)
	if err != nil {
		// Connection or statement error
		return nil, err
	}
	defer rows.Close()

	teams := []*Team{}
	for rows.Next() {
		t := Team{}
		err := rows.Scan(&t.id, &t.name, &t.division, &t.logoURL)
		if err != nil {
			// Row parsing error
			return nil, err
		}
		teams = append(teams, &t)
		if len(teams) > 1 {
			return teams[0], fmt.Errorf("expected only one team to match the criteria but found %d.", len(teams))
		}
	}
	err = rows.Err()
	if err != nil {
		// Errors within rows
		return nil, err
	}
	rows.Close()
	return teams[0], nil
}

//-----------------------------------------------//
// Division
//-----------------------------------------------//
type Division struct {
	id			uint32
	name		string
}

func (d *Division) Save() (id uint32, err error) {
	if d.id > 0 {
		return d.Update()
	} else {
		return d.Create()
	}
}

func (d *Division) Create() (id uint32, err error) {
	id = 0
	err = database.DB.QueryRow("INSERT INTO divisions (name) VALUES (?) RETURNING division_id", d.name).Scan(&id)
	return id, err
}

func (d *Division) Update() (id uint32, err error) {
	id = 0
	if d.id > 0 {
		// try updating using the ID
		err = database.DB.QueryRow("UPDATE divisions SET name = ? WHERE division_id = ? RETURNING division_id", d.name, d.id).Scan(&id)
	} else {
		// try updating using the name
		// this is obsolete while it's only name and ID
	}
	return id, err
}

//-----------------------------------------------//
// Player
//-----------------------------------------------//
type Player struct {
	id			uint32
	number		uint
	name		string
	teams		[]string
	position	string
}

//-----------------------------------------------//
// Goalie
//-----------------------------------------------//
type Goalie struct {
	id			uint32
	number		uint
	name		string
	teams		[]string
	shots		uint
	saves		uint
	mins		uint
}
