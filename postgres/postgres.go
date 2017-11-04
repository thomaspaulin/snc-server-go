package postgres

import (
	"database/sql"
	"github.com/thomaspaulin/snc-server-go/snc"
	"log"
)


//--------------------------------------------------------------------------------------------------------------------//
// Divisions
//--------------------------------------------------------------------------------------------------------------------//
// DivisionService represents the PostgreSQL implementation of snc.DivisionService
type DivisionService struct {
	DB *sql.DB
}

func (ds *DivisionService) CreateDivision(d *snc.Division) error {
	// the returned ID is not being used for the time being
	id := 0
	err := ds.DB.QueryRow(`
		INSERT INTO divisions
			(name)
		VALUES
			($1)
		RETURNING division_id`, d.Name).Scan(&id)
	return err
}

func (ds *DivisionService) Division(id int) (*snc.Division, error) {
	d := snc.Division{ID: uint32(id)}
	err := ds.DB.QueryRow(`
	SELECT
		division_id, name
	FROM divisions
	WHERE division_id = $1 AND deleted IS FALSE`, id).Scan(&d.ID, &d.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &d, nil
}

func (ds *DivisionService) Divisions() ([]*snc.Division, error) {
	rows, err := ds.DB.Query(`
	SELECT
  		division_id, name
	FROM divisions
	WHERE deleted IS FALSE`)
	if err != nil {
		// Connection or statement error
		return nil, err
	}
	defer rows.Close()

	divs := []*snc.Division{}
	for rows.Next() {
		d := &snc.Division{}
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

func (ds *DivisionService) UpdateDivision(d *snc.Division) error {
	err := ds.DB.QueryRow(`
		UPDATE divisions SET
			name = $1
		WHERE
			division_id = $2 AND deleted IS FALSE
		RETURNING division_id`, d.Name, d.ID).Scan()
	if err != nil {
		log.Println(err.Error());
	}
	return err
}

func (ds *DivisionService) DeleteDivision(id int) error {
	err := ds.DB.QueryRow(`
		UPDATE divisions SET
			deleted = TRUE
		WHERE
			division_id = $1`).Scan()
	if err != nil {
		log.Println(err.Error());
	}
	return err
}


//--------------------------------------------------------------------------------------------------------------------//
// Rinks
//--------------------------------------------------------------------------------------------------------------------//
// RinkService represents the PostgreSQL implementation of snc.RinkService
type RinkService struct {
	DB *sql.DB
}

func (rs *RinkService) CreateRink(r *snc.Rink) error {
	err := rs.DB.QueryRow(`
		INSERT INTO rinks
			(name)
		VALUES
			($1)
		RETURNING rink_id`, r.Name).Scan()
	if err != nil {
		log.Println(err.Error())
	}
	return err
}

func (rs *RinkService) Rink(id int) (*snc.Rink, error) {
	r := snc.Rink{ID: uint32(id)}
	err := rs.DB.QueryRow(`
	SELECT
		rink_id, name
	FROM rinks
	WHERE rink_id = $1 AND deleted IS FALSE`, id).Scan(&r.ID, &r.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &r, nil
}

func (rs *RinkService) Rinks() ([]*snc.Rink, error) {
	rows, err := rs.DB.Query(`
	SELECT
  		rink_id, name
	FROM rinks
	WHERE deleted IS FALSE`)
	if err != nil {
		// Connection or statement error
		return nil, err
	}
	defer rows.Close()

	rinks := []*snc.Rink{}
	for rows.Next() {
		r := &snc.Rink{}
		err := rows.Scan(&r.ID, &r.Name)
		// err here is the row.Scan(...) error. It shadows the err from outside the loop, and does not overwrite
		// skip the rink if there's an error but log it anyway
		if err != nil {
			// probably the schema is wrong or the row is bad and so the database needs inspecting
			// later on this might want to be changed to pass through and list the IDs of the bad rows
			log.Printf("postgres: error encountered when scanning a row. Row will be logged but here is the error: %s\n", err.Error())
		}
		rinks = append(rinks, r)
	}
	err = rows.Err()
	if err != nil {
		// Errors within rows
		return nil, err
	}
	rows.Close()
	return rinks, nil
}

func (rs *RinkService) UpdateRink(r *snc.Rink) error {
	err := rs.DB.QueryRow(`
	UPDATE rinks
	SET
		name = $1
	WHERE
		rink_id = $2 AND deleted IS FALSE`, r.Name, r.ID).Scan()
	if err != nil {
		// in future when there are more columns I'd use the name here to uniquely identify rinks and update the other
		// columns but at present it's a bit pointless looking up using name then updating name (ID should be fixed)
		log.Println(err.Error())
	}
	return err
}

func (rs *RinkService) DeleteRink(id int) error {
	err := rs.DB.QueryRow(`
		UPDATE rinks SET
			deleted = TRUE
		WHERE
			rink_id = $1`).Scan()
	if err != nil {
		log.Println(err.Error());
	}
	return err
}


//--------------------------------------------------------------------------------------------------------------------//
// Teams
//--------------------------------------------------------------------------------------------------------------------//
// TeamService represents the PostgreSQL implementation of snc.TeamService
type TeamService struct {
	DB *sql.DB
}

func (ts *TeamService) CreateTeam(t *snc.Team) error {
	id := 0
	// TODO save the division ID rather than just the 0. Be it through a query or having the division as a field on the team struct
	err := ts.DB.QueryRow(`
	INSERT INTO teams
		(name, division_id, logo_url)
	VALUES
  		($1, $2, $3)
	RETURNING team_id`, t.Name, 0, t.LogoURL).Scan(&id)
	if err != nil {
		log.Println(err.Error())
	}
	return err
}

func (ts *TeamService) Team(id int) (*snc.Team, error) {
	t := snc.Team{ID: uint32(id)}
	err := ts.DB.QueryRow(`
	SELECT
  		teams.name     AS  team_name,
  		divisions.name AS  div_name,
  		teams.logo_url
	FROM (teams
		JOIN divisions ON teams.division_id = divisions.division_id
	 )
	 WHERE teams.team_id = $1 AND teams.deleted IS FALSE`, id).Scan(&t.Name, &t.Division, &t.LogoURL)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &t, nil
}

func (ts *TeamService) Teams() ([]*snc.Team, error) {
	// TODO create a new division if there isn't one?
	rows, err := ts.DB.Query(`
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

	teams := []*snc.Team{}
	for rows.Next() {
		t := snc.Team{}
		err := rows.Scan(&t.ID, &t.Name, &t.Division, &t.LogoURL)
		if err != nil {
			// Row parsing error. Pass over it but log it too
			log.Println(err.Error())
		}
		teams = append(teams, &t)
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

func (ts *TeamService) UpdateTeam(t *snc.Team) error {
	err := ts.DB.QueryRow(`
		UPDATE teams SET
			name = $1
		WHERE
			team_id = $2 AND deleted IS FALSE`, t.Name, t.ID).Scan()
	if err != nil {
		log.Println(err.Error());
	}
	return err
}

func (ts *TeamService) DeleteTeam(id int) error {
	err := ts.DB.QueryRow(`
		UPDATE teams SET
			deleted = TRUE
		WHERE
			team_id = $1`).Scan()
	if err != nil {
		log.Println(err.Error());
	}
	return err
}