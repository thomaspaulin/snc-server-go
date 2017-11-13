package snc

import (
	"errors"
	"database/sql"
	"log"
	"strings"
)

//--------------------------------------------------------------------------------------------------------------------//
// Team
//--------------------------------------------------------------------------------------------------------------------//
type Team struct {
	ID			uint	`json:"id"`
	Name		string	`json:"name"`
	Division	string	`json:"division"`
	LogoURL		string	`json:"logoURL"`
	Players		[]Player`json:"players,omitempty"`
}

var ErrMultipleTeams = errors.New("snc: expected only one team but got multiple")

func CreateTeam(t Team, DB *sql.DB) error {
	id := 0
	// TODO save the division ID rather than just the 0. Be it through a query or having the division as a field on the team struct
	err := DB.QueryRow(`
	INSERT INTO teams
		(name, division_id, logo_url)
	VALUES
  		($1, $2, $3)
	RETURNING team_id`, strings.ToLower(t.Name), 0, t.LogoURL).Scan(&id)
	if err != nil {
		log.Println(err.Error())
	}
	return err
}

func FetchTeam(id uint, DB *sql.DB) (Team, error) {
	t := Team{ID: id}
	err := DB.QueryRow(`
	SELECT
  		teams.name     AS  team_name,
  		divisions.name AS  div_name,
  		teams.logo_url
	FROM (teams
		JOIN divisions ON teams.division_id = divisions.division_id
	 )
	 WHERE teams.team_id = $1 AND teams.deleted IS FALSE`, id).Scan(&t.Name, &t.Division, &t.LogoURL)
	if err == sql.ErrNoRows {
		return t, nil
	} else if err != nil {
		log.Println(err.Error())
		return Team{Name: "Unknown"}, err
	}
	return t, nil
}

func FetchTeams(DB *sql.DB) ([]Team, error) {
	// TODO create a new division if there isn't one?
	rows, err := DB.Query(`
	SELECT
  		teams.team_id  AS team_id,
  		teams.name     AS team_name,
  		divisions.name AS div_name,
  		teams.logo_url
	FROM (teams
		JOIN divisions
      	ON teams.division_id = divisions.division_id
	)
	WHERE teams.deleted IS FALSE
    ORDER BY team_id`)
	if err != nil {
		log.Println(err.Error())
		// Connection or statement error
		return nil, err
	}
	defer rows.Close()

	teams := []Team{}
	for rows.Next() {
		t := Team{}
		err := rows.Scan(&t.ID, &t.Name, &t.Division, &t.LogoURL)
		if err != nil {
			// Row parsing error. Pass over it but log it too
			log.Println(err.Error())
		}
		teams = append(teams, t)
	}
	err = rows.Err()
	if err != nil {
		// Errors within rows
		log.Println(err.Error())
		return nil, err
	}
	rows.Close()
	return teams, nil
}

func TeamCalled(name string, DB *sql.DB) (Team, error) {
	t := Team{Name: name}
	err := DB.QueryRow(`
	SELECT
  		teams.team_id  AS  team_id,
  		divisions.name AS  div_name,
  		teams.logo_url
	FROM (teams
		JOIN divisions ON teams.division_id = divisions.division_id
	 )
	 WHERE teams.name = $1 AND teams.deleted IS FALSE`, strings.ToLower(name)).Scan(&t.ID, &t.Division, &t.LogoURL)
	if err == sql.ErrNoRows {
		return t, nil
	} else if err != nil {
		log.Println(err.Error())
		return Team{}, err
	}
	return t, nil
}

func UpdateTeam(t *Team, DB *sql.DB) error {
	id := 0
	err := DB.QueryRow(`
		UPDATE teams SET
			name = $1
		WHERE
			team_id = $2 AND deleted IS FALSE
		RETURNING team_id`, strings.ToLower(t.Name), t.ID).Scan(&id)
	if err != nil {
		log.Println(err.Error());
	}
	return err
}

func DeleteTeam(id int, DB *sql.DB) error {
	// Deleting a team should NOT delete a division because divisions don't have dependencies on teams
	deleted := false
	err := DB.QueryRow(`
		UPDATE teams SET
			deleted = TRUE
		WHERE
			team_id = $1
		RETURNING deleted`, id).Scan(&deleted)
	if err != nil {
		log.Println(err.Error());
	}
	return err
}

//--------------------------------------------------------------------------------------------------------------------//
// Division
//--------------------------------------------------------------------------------------------------------------------//
type Division struct {
	ID			uint	`json:"id"`
	Name		string	`json:"name"`
}

func CreateDivision(d Division, DB *sql.DB) error {
	// the returned ID is not being used for the time being
	id := 0
	err := DB.QueryRow(`
		INSERT INTO divisions
			(name)
		VALUES
			($1)
		RETURNING division_id`, d.Name).Scan(&id)
	return err
}

func FetchDivision(id uint, DB *sql.DB) (Division, error) {
	d := Division{ID: id}
	err := DB.QueryRow(`
	SELECT
		division_id, name
	FROM divisions
	WHERE division_id = $1 AND deleted IS FALSE`, id).Scan(&d.ID, &d.Name)
	if err == sql.ErrNoRows {
		return d, nil
	} else if err != nil {
		return Division{Name: "Unknown"}, err
	}
	return d, nil
}

func FetchDivisions(DB *sql.DB) ([]Division, error) {
	rows, err := DB.Query(`
	SELECT
  		division_id, name
	FROM divisions
	WHERE deleted IS FALSE`)
	if err != nil {
		// Connection or statement error
		return nil, err
	}
	defer rows.Close()

	divs := []Division{}
	for rows.Next() {
		d := Division{}
		err := rows.Scan(&d.ID, &d.Name)
		// skip the division if there's an error but log it anyway
		if err != nil {
			log.Printf("postgres: error encountered when scanning a row. Row will be logged but here is the error: %s\n", err.Error())
		}
		divs = append(divs, d)
	}
	err = rows.Err()
	if err != nil {
		// Errors within rows
		return nil, err
	}
	rows.Close()
	return divs, nil
}

func UpdateDivision(d Division, DB *sql.DB) error {
	id := 0
	err := DB.QueryRow(`
		UPDATE divisions SET
			name = $1
		WHERE
			division_id = $2 AND deleted IS FALSE
		RETURNING division_id`, d.Name, d.ID).Scan(&id)
	if err != nil {
		log.Println(err.Error());
	}
	return err
}

func DeleteDivision(id int, DB *sql.DB) error {
	// todo deleting a division should delete all the teams in the division because teams are tied to divisions
	deleted := false
	err := DB.QueryRow(`
		UPDATE divisions SET
			deleted = TRUE
		WHERE
			division_id = $1
		RETURNING deleted`).Scan(&deleted)
	if err != nil {
		log.Println(err.Error());
	}
	return err
}

//--------------------------------------------------------------------------------------------------------------------//
// Player (Goalies included)
//--------------------------------------------------------------------------------------------------------------------//
type Player struct {
	ID			uint		`json:"id"`
	Number		uint		`json:"number"`
	Name		string		`json:"name"`
	//Teams		[]string	`json:"teams"`
	Position	string		`json:"position"`
}
