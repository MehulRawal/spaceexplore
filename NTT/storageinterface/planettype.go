package models

import (
	"database/sql"
	"log"
)

type PlanetType struct {
	ID          int    `json:"id"`
	PlanetType  string `json:"planet_type"`
	Description string `json:"description"`
}

func GetPlanetTypes(db *sql.DB) ([]PlanetType, error) {

	selectPlanetTypeQuery := `SELECT id, planet_type, description FROM planet_types`
	rows, err := db.Query(selectPlanetTypeQuery)
	if err != nil {
		log.Println("models.GetPlanetTypes error in selectPlanetTypeQuery : ", err.Error())
		return nil, err
	}
	defer rows.Close()

	var planetTypes []PlanetType
	for rows.Next() {
		var pt PlanetType
		if err := rows.Scan(&pt.ID, &pt.PlanetType, &pt.Description); err != nil {
			log.Println("models.GetPlanetTypes error while scanning selectPlanetTypeQuery : ", err.Error())
			return nil, err
		}
		planetTypes = append(planetTypes, pt)
	}

	return planetTypes, nil
}

func CheckIfPlanetTypeExist(db *sql.DB, planetType string) (bool, error) {

	var count int
	selectCountPlanetTypeQuery := `Select count(*) from planet_types where planet_type = ?`
	err := db.QueryRow(selectCountPlanetTypeQuery, planetType).Scan(count)
	if err != nil && err != sql.ErrNoRows {
		log.Println("models.CheckIfPlanetTypeExist err in selectCountPlanetTypeQuery : ", err.Error())
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func CheckIfPlanetTypeByIDExist(db *sql.DB, id int) (bool, error) {

	var count int
	selectCountPlanetTypeByIDQuery := `Select count(*) from planet_types where id = ?`
	err := db.QueryRow(selectCountPlanetTypeByIDQuery, id).Scan(count)
	if err != nil && err != sql.ErrNoRows {
		log.Println("models.CheckIfPlanetTypeByIDExist err in selectCountPlanetTypeByIDQuery : ", err.Error())
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func InsertPlanetTypes(db *sql.DB, request PlanetType) (int, error) {

	insertPlanetTypeQuery := `INSERT INTO  planet_types(planet_type, description) VALUES(?, ?)`

	result, err := db.Exec(insertPlanetTypeQuery, request.PlanetType, request.Description)
	if err != nil {
		log.Println("models.InsertPlanetTypes err in insertPlanetTypeQuery : ", err.Error())
		return 0, err
	}

	var id int
	if result != nil {
		lastId, err := result.LastInsertId()
		if err != nil {
			log.Println("models.InsertPlanetTypes err in LastInsertId : ", err.Error())
			return 0, err
		}
		id = int(lastId)
	}

	return id, nil
}
