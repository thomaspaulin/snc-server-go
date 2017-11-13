package snc

import (
	"log"
	"database/sql"
)

const (
	Avondale = "Avondale"
	Botany = "Botany"
)

type Rink struct {
	ID		uint	`json:"id"`
	Name	string	`json:"name"`
}


func CreateRink(r Rink, DB *sql.DB) error {
	id := 0
	err := DB.QueryRow(`
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

func FetchRink(id uint, DB *sql.DB) (Rink, error) {
	r := Rink{ID: id}
	err := DB.QueryRow(`
	SELECT
		rink_id, name
	FROM rinks
	WHERE rink_id = $1 AND deleted IS FALSE`, id).Scan(&r.ID, &r.Name)
	if err == sql.ErrNoRows {
		return Rink{Name: "Unknown"}, nil
	} else if err != nil {
		return r, err
	}
	return r, nil
}

func FetchRinks(DB *sql.DB) ([]Rink, error) {
	rows, err := DB.Query(`
	SELECT
  		rink_id, name
	FROM rinks
	WHERE deleted IS FALSE`)
	if err != nil {
		// Connection or statement error
		return nil, err
	}
	defer rows.Close()

	rinks := []Rink{}
	for rows.Next() {
		r := Rink{Name: "Unknown"}
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

func UpdateRink(r Rink, DB *sql.DB) error {
	id := 0
	err := DB.QueryRow(`
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

func DeleteRink(id int, DB *sql.DB) error {
	deleted := false
	err := DB.QueryRow(`
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