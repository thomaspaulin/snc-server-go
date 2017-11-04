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
	id := 0
	err := ds.DB.QueryRow(`
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

func (ds *DivisionService) DeleteDivision(id int) error {
	// todo deleting a division should delete all the teams in the division because teams are tied to divisions
	deleted := false
	err := ds.DB.QueryRow(`
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
// Matches
//--------------------------------------------------------------------------------------------------------------------//
// MatchServices represents the PostgreSQL implementation of snc.MatchService
type MatchService struct {
	DB *sql.DB
}

// todo match divisions

//func (ms *MatchService) CreateMatch(m *snc.Match) error {
//	id := 0
//	err := ms.DB.QueryRow(`
//	INSERT INTO matches
//	 	(start, season, status, away_id, home_id, away_score, home_score, rink_id)
//	VALUES
//		($1, $2, $3, $4, $5, 0, 0, $6)
//	RETURNING match_id`, m.Start, m.Season, m.Status, awayID, homeID, m.AwayScore, m.HomeScore, rinkID).Scan(&id)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	return err
//}
//
//func (ms *MatchService) Match(id int) (*snc.Match, error) {
//	m := snc.Match{}
//	err := ms.DB.QueryRow(`
//	SELECT
//		m.match_id,
//		m.start,
//		m.season,
//		m.status,
//		away.name,
//		home.name,
//		m.away_score,
//		m.home_score,
//		rink.name
//	FROM (matches AS m
//		INNER JOIN teams AS home
//			ON home.team_id = m.home_id
//		INNER JOIN teams AS away
//			ON away.team_id = m.away_id
//		INNER JOIN rinks AS rink
//			ON rink.rink_id = m.rink_id)
//	WHERE m.match_id = $1 AND m.deleted IS FALSE`, id).
//		Scan(&m.ID, &m.Start, &m.Season, &m.Status, &m.Away, &m.Home, &m.AwayScore, &m.HomeScore, &m.Rink)
//	if err == sql.ErrNoRows {
//		return nil, nil
//	} else if err != nil {
//		return nil, err
//	}
//	return &m, nil
//}
//
//func (ms *MatchService) Matches() ([]*snc.Match, error) {
//	rows, err := ms.DB.Query(`
//	SELECT
//		m.match_id,
//		m.start,
//		m.season,
//		m.status,
//		away.name,
//		home.name,
//		m.away_score,
//		m.home_score,
//		rink.name
//	FROM (matches AS m
//		INNER JOIN teams AS home
//			ON home.team_id = m.home_id
//		INNER JOIN teams AS away
//			ON away.team_id = m.away_id
//		INNER JOIN rinks AS rink
//			ON rink.rink_id = m.rink_id)
//	WHERE m.deleted IS FALSE
//	ORDER BY m.start DESC`)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	matches := make([]*snc.Match, 0)
//	for rows.Next() {
//		m := snc.Match{}
//		err := rows.Scan(&m.ID, &m.Start, &m.Season, &m.Status, &m.Away, &m.Home, &m.AwayScore, &m.HomeScore, &m.Rink)
//		// err here is the row.Scan(...) error. It shadows the err from outside the loop, and does not overwrite
//		if err != nil {
//			// probably the schema is wrong or the row is bad and so the database needs inspecting
//			// later on this might want to be changed to pass through and list the IDs of the bad rows
//			return nil, err
//		}
//		matches = append(matches, &m)
//	}
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//	return matches, nil
//}
//
//func (ms *MatchService) UpdateMatch(m *snc.Match) error {
//
//}
//
//func (ms *MatchService) DeleteMatch(id int) error {
//
//}


//--------------------------------------------------------------------------------------------------------------------//
// Rinks
//--------------------------------------------------------------------------------------------------------------------//
// RinkService represents the PostgreSQL implementation of snc.RinkService
type RinkService struct {
	DB *sql.DB
}

func (rs *RinkService) CreateRink(r *snc.Rink) error {
	id := 0
	err := rs.DB.QueryRow(`
		INSERT INTO rinks
			(name)
		VALUES
			($1)
		RETURNING rink_id`, r.Name).Scan(&id)
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
	id := 0
	err := rs.DB.QueryRow(`
	UPDATE rinks
	SET
		name = $1
	WHERE
		rink_id = $2 AND deleted IS FALSE
	RETURNING rink_id`, r.Name, r.ID).Scan(&id)
	if err != nil {
		// in future when there are more columns I'd use the name here to uniquely identify rinks and update the other
		// columns but at present it's a bit pointless looking up using name then updating name (ID should be fixed)
		log.Println(err.Error())
	}
	return err
}

func (rs *RinkService) DeleteRink(id int) error {
	deleted := false
	err := rs.DB.QueryRow(`
		UPDATE rinks SET
			deleted = TRUE
		WHERE
			rink_id = $1
		RETURNING deleted`, id).Scan(&deleted)
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
	id := 0
	err := ts.DB.QueryRow(`
		UPDATE teams SET
			name = $1
		WHERE
			team_id = $2 AND deleted IS FALSE
		RETURNING team_id`, t.Name, t.ID).Scan(&id)
	if err != nil {
		log.Println(err.Error());
	}
	return err
}

func (ts *TeamService) DeleteTeam(id int) error {
	// Deleting a team should NOT delete a division because divisions don't have dependencies on teams
	deleted := false
	err := ts.DB.QueryRow(`
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