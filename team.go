package main

import (
	"database/sql"
	"fmt"
	"log"
)

//-----------------------------------------------//
// Team
//-----------------------------------------------//
type Team struct {
	ID			uint32	`json:"id"`
	Name		string	`json:"name"`
	Division	string	`json:"division"`
	LogoURL		string	`json:"logoURL"`
}

func (t *Team) Create(db *sql.DB) (id uint32, err error) {
	d := Division{Name: t.Division}
	divID, err := d.Save(db)
	id = 0
	err = db.QueryRow(`
	INSERT INTO teams
		(name, division_id, logo_url)
	VALUES
  		($1, $2, $3)
	RETURNING team_id`, t.Name, divID, t.LogoURL).Scan(&id)
	if err != nil {
		log.Println(err.Error())
	}
	return id, nil
}

func FetchTeamByID(db *sql.DB, id uint32) (*Team, error) {
	t := Team{ID: id}
	err := db.QueryRow(`
	SELECT
  		teams.name     AS  team_name,
  		divisions.name AS  div_name,
  		teams.logo_url
	FROM (teams
		JOIN divisions ON teams.division_id = divisions.division_id
	 )
	 WHERE team_id = $1`, id).Scan(&t.Name, &t.Division, &t.LogoURL)
	if err == sql.ErrNoRows {
		return &Team{}, nil
	} else if err != nil {
		return nil, err
	}
	return &t, nil
}

func FetchTeam(db *sql.DB, teamName string, divName string) (*Team, error) {
	rows, err := db.Query(`
	SELECT
  		teams.team_id  AS team_id,
  		teams.name     AS team_name,
  		divisions.name AS div_name,
  		teams.logo_url
	FROM (teams
		JOIN divisions
      	ON teams.division_id = divisions.division_id
	)
	WHERE teams.name = $1
    AND divisions.name = $2`, teamName, divName)
	if err != nil {
		// Connection or statement error
		return nil, err
	}
	defer rows.Close()

	teams := []*Team{}
	for rows.Next() {
		t := Team{}
		err := rows.Scan(&t.ID, &t.Name, &t.Division, &t.LogoURL)
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
	ID			uint32	`json:"id"`
	Name		string	`json:"name"`
}

func (d *Division) Save(db *sql.DB) (id uint32, err error) {
	if d.ID > 0 {
		return d.Update(db)
	} else {
		return d.Create(db)
	}
}

func (d *Division) Create(db *sql.DB) (id uint32, err error) {
	id = 0
	err = db.QueryRow(`
		INSERT INTO divisions
			(name)
		VALUES
			($1)
		RETURNING division_id`, d.Name).Scan(&id)
	return id, err
}

func (d *Division) Update(db *sql.DB) (id uint32, err error) {
	id = 0
	if d.ID > 0 {
		// try updating using the ID
		err = db.QueryRow(`
			UPDATE divisions
			SET name = $1
			WHERE division_id = $2
			RETURNING division_id`, d.Name, d.ID).Scan(&id)
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
	ID			uint32		`json:"id"`
	Number		uint		`json:"number"`
	Name		string		`json:"name"`
	Teams		[]string	`json:"teams"`
	Position	string		`json:"position"`
}

//-----------------------------------------------//
// Goalie
//-----------------------------------------------//
type Goalie struct {
	ID			uint32		`json:"id"`
	Number		uint		`json:"number"`
	Name		string		`json:"name"`
	Teams		[]string	`json:"teams"`
	Shots		uint		`json:"shots"`
	Saves		uint		`json:"saves"`
	Mins		uint		`json:"mins"`
}
