package snc

import (
	"database/sql"
	"log"
)

const (
	Avondale = "Avondale"
	Botany = "Botany"
)

type Rink struct {
	ID		uint32
	Name	string
}

func (r Rink) Create(DB *sql.DB) (id uint32, err error) {
	err = DB.QueryRow(`
	INSERT INTO rinks
		(name)
	VALUES
		($1)
	RETURNING rink_id`, r.ID).Scan(&id)
	if err != nil {
		log.Println(err.Error())
	}
	return id, err
}

func FetchRinks(DB *sql.DB) ([]Rink, error) {
	rows, err := DB.Query(`
	SELECT
		rink_id,
		name
	FROM rinks
	ORDER BY name DESC`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	rinks := make([]Rink, 0)
	for rows.Next() {
		r := Rink{}
		err := rows.Scan(&r.ID, &r.Name)
		// err here is the row.Scan(...) error. It shadows the err from outside the loop, and does not overwrite
		if err != nil {
			// probably the schema is wrong or the row is bad and so the database needs inspecting
			// later on this might want to be changed to pass through and list the IDs of the bad rows
			return nil, err
		}
		rinks = append(rinks, r)
	}
	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return rinks, nil
}

func FetchRinkByID(DB *sql.DB, rinkID uint32) (r Rink, err error) {
	r.ID = rinkID
	err = DB.QueryRow(`
	SELECT
		rink_id,
		name
	FROM rinks
	WHERE rink_id = $1`, rinkID).Scan(&r.ID, &r.Name)
	if err != nil {
		log.Println(err.Error())
	}
	return r, err
}

func FetchRink(DB *sql.DB, rinkName string) (r Rink, err error) {
	err = DB.QueryRow(`
	SELECT
		rink_id,
		name
	FROM rinks
	WHERE name = $1`, rinkName).Scan(&r.ID, &r.Name)
	if err != nil {
		log.Println(err.Error())
	}
	return r, err
}

func (r Rink) Update(DB *sql.DB) (id uint32, err error) {
	err = DB.QueryRow(`
	UPDATE rinks
	SET
		name = $1
	WHERE
		rink_id = $2
	RETURNING rink_id`, r.Name, r.ID).Scan(&id)
	if err != nil {
		// in future when there are more columns I'd use the name here to uniquely identify rinks and update the other
		// columns but at present it's a bit pointless looking up using name then updating name (ID should be fixed)
		log.Println(err.Error())
	}
	return id, err
}
