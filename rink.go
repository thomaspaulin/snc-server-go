package main

import "database/sql"

const (
	Avondale = "Avondale"
	Botany = "Botany"
)

var UnknownRink = Rink{0, "Unknown"}

type Rink struct {
	ID		uint32
	Name	string
}

func FetchRinks(db *sql.DB) ([]*Rink, error) {
	rows, err := db.Query(`
	SELECT
		rink_id,
		name
	FROM rinks
	ORDER BY name DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rinks := make([]*Rink, 0)
	for rows.Next() {
		r := Rink{}
		err := rows.Scan(&r.ID, &r.Name)
		// err here is the row.Scan(...) error. It shadows the err from outside the loop, and does not overwrite
		if err != nil {
			// probably the schema is wrong or the row is bad and so the database needs inspecting
			// later on this might want to be changed to pass through and list the IDs of the bad rows
			return nil, err
		}
		rinks = append(rinks, &r)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rinks, nil
}

func FetchRinkByID(db *sql.DB, rinkID uint32) (*Rink, error) {
	r := Rink{ID: rinkID}
	err := db.QueryRow(`
	SELECT
		rink_id,
		name
	FROM rinks
	WHERE rink_id = $1`, rinkID).Scan(&r.ID, &r.Name)
	if err != nil {
		return nil, err
	} else {
		return &r, nil
	}
}

func FetchRink(db *sql.DB, rinkName string) (*Rink, error) {
	r := Rink{}
	err := db.QueryRow(`
	SELECT
		rink_id,
		name
	FROM rinks
	WHERE name = $1`, rinkName).Scan(&r.ID, &r.Name)
	if err != nil {
		return &Rink{}, err
	}
	return &r, nil
}