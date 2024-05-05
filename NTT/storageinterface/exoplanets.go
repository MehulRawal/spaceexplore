package models

import (
	"database/sql"
	"log"
)

type PlanetType struct {
	ID          int    `json:"id"`
	PlanetType  string `json:"name"`
	Description string `json:"email"`
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
