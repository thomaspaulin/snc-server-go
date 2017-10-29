package main

import (
	"github.com/thomaspaulin/snc-server-go/database"
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

func FetchTeamByID(id uint32) (*Team, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	t := &Team{id: id}
	err = tx.QueryRow("SELECT teams.name AS team_name, " +
								"divisions.name AS div_name, " +
								"teams.logo_url " +
							"FROM teams " +
							"JOIN divisions " +
								"ON teams.division_id = divisions.division_id " +
							"WHERE team_id = ?", id).Scan(&t.id, &t.division, &t.logoURL)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return t, nil
}

func FetchTeam(teamName string, divName string) (*Team, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	t := &Team{name: teamName, division: divName}

	rows, err := tx.Query("SELECT teams.team_id AS team_id, " +
									"teams.name AS team_name, " +
									"divisions.name AS div_name, " +
									"teams.logo_url " +
								"FROM teams " +
								"JOIN divisions " +
									"ON teams.division_id = divisions.division_id " +
								"WHERE teams.name = ? " +
									"AND divisions.name = ?", teamName, divName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&t.id, &t.name, &t.division, &t.logoURL)
		if err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return t, nil
}

//-----------------------------------------------//
// Division
//-----------------------------------------------//
type Division struct {
	id			uint32
	name		string
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
