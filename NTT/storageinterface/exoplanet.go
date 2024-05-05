package models

import (
	"database/sql"
	"log"
	"strings"
)

// currently assuming Mass in KG and Radius in Kms
type ExoPlanet struct {
	ID                int     `json:"id"`
	Name              string  `json:"string"`
	Description       string  `json:"description"`
	DistanceFromEarth int     `json:"distance_from_earth"`
	Radius            float64 `json:"radius"`
	Mass              float64 `json:"mass"`
	PlanetTypeID      int     `json:"planet_type_id"`
	PlanetType        string  `json:"planet_type"`
}

type ExoplanetFilters struct {
	Radius float64 `json:"radius"`
	Mass   float64 `json:"mass"`
}

func CheckIfNameAlreadyExist(db *sql.DB, name string) (bool, error) {

	var count int
	selectCountExoplanetByNameQuery := `Select count(*) from exoplanets where name = ?`
	err := db.QueryRow(selectCountExoplanetByNameQuery, name).Scan(count)
	if err != nil && err != sql.ErrNoRows {
		log.Println("models.CheckIfNameAlreadyExist err in selectCountExoplanetByNameQuery : ", err.Error())
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func CheckIfExoplanetExist(db *sql.DB, id int) (bool, error) {

	var count int
	selectCountExoplanetByIDQuery := `Select count(*) from exoplanets where id = ?`
	err := db.QueryRow(selectCountExoplanetByIDQuery, id).Scan(count)
	if err != nil && err != sql.ErrNoRows {
		log.Println("models.CheckIfExoplanetExist err in selectCountExoplanetByIDQuery : ", err.Error())
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func CheckIfNameAlreadyExist(db *sql.DB, name string) (bool, error) {

	var count int
	selectCountExoplanetByNameQuery := `Select count(*) from exoplanets where name = ?`
	err := db.QueryRow(selectCountExoplanetByNameQuery, name).Scan(count)
	if err != nil && err != sql.ErrNoRows {
		log.Println("models.CheckIfNameAlreadyExist err in selectCountExoplanetByNameQuery : ", err.Error())
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func InsertExoplanet(db *sql.DB, request *ExoPlanet) (int, error) {

	insertExoplanetQuery := `INSERT INTO exoplanets(name, description, distance_from_earth, radius, mass, planet_type_id) VALUES(?, ?, ?, ?, ?, ?)`

	result, err := db.Exec(insertExoplanetQuery, request.Name, request.Description, request.DistanceFromEarth, request.Radius, request.Mass, request.PlanetTypeID)
	if err != nil {
		log.Println("models.InsertExoplanet error in insertExoplanetQuery : ", err.Error())
		return 0, err
	}

	var id int

	if result != nil {
		lastId, err := result.LastInsertId()
		if err != nil {
			log.Println("models.InsertExoplanet err in LastInsertId : ", err.Error())
			return 0, err
		}
		id = int(lastId)
	}

	return id, nil
}

func UpdateExoplanet(db *sql.DB, id int, request *ExoPlanet) error {
	updateExoplanetQuery := `Update exoplanets set  name = ?, description = ?, distance_from_earth = ?,  radius = ?, mass = ?, planet_type_id = ? where id = ?`

	_, err := db.Exec(updateExoplanetQuery, request.Name, request.Description, request.DistanceFromEarth, request.Radius, request.Mass, request.PlanetTypeID, id)
	if err != nil {
		log.Println("models.UpdateExoplanet error in updateExoplanetQuery : ", err.Error())
		return err
	}

	return nil
}

func GetExoplanets(db *sql.DB, filters *ExoplanetFilters) ([]ExoPlanet, error) {

	selectExoplanetQuery := `select e.id, e.name, COALESCE(e.description,''), COALESCE(e.distance_from_earth, 0), COALESCE(e.radius, 0), COALESCE(e.mass, 0), pt.id, COALESCE(pt.planet_type, '')  from exoplanets e JOIN planet_types pt on e.planet_type_id = pt.id`

	values := make([]interface{}, 0)
	var whereClause []string
	if filters != nil {

		if filters.Radius != 0 {
			whereClause = append(whereClause, " e.radius = ? ")
			values = append(values, filters.Radius)
		}
		if filters.Mass != 0 {
			whereClause = append(whereClause, " e.mass = ? ")
			values = append(values, filters.Mass)
		}

	}

	if len(whereClause) > 0 {
		selectExoplanetQuery = selectExoplanetQuery + " where " + strings.Join(whereClause, " AND ")
	}

	rows, err := db.Query(selectExoplanetQuery, values...)
	if err != nil {
		log.Println("models.GetExoplanets error in selectExoplanetQuery : ", err.Error())
		return nil, err
	}
	defer rows.Close()

	var exoPlanets []ExoPlanet
	for rows.Next() {
		var exoplanet ExoPlanet
		err := rows.Scan(&exoplanet.ID, &exoplanet.Name, &exoplanet.Description, &exoplanet.DistanceFromEarth, &exoplanet.Radius, &exoplanet.Mass, &exoplanet.PlanetTypeID, &exoplanet.PlanetType)
		if err != nil {
			log.Println("models.GetExoplanets error while scanning selectExoplanetQuery : ", err.Error())
			return nil, err
		}
		exoPlanets = append(exoPlanets, exoplanet)
	}

	return exoPlanets, nil
}

func GetExoplanetByID(db *sql.DB, id int) (*ExoPlanet, error) {
	selectExoplanetByIDQuery := `select e.id, e.name, COALESCE(e.description,''), COALESCE(e.distance_from_earth, 0), COALESCE(e.radius, 0), COALESCE(e.mass, 0), pt.id, COALESCE(pt.planet_type, '')  from exoplanets e JOIN planet_types pt on e.planet_type_id = pt.id where e.id = ?`

	exoplanet := ExoPlanet{}
	err := db.QueryRow(selectExoplanetByIDQuery, id).Scan(&exoplanet.ID, &exoplanet.Name, &exoplanet.Description, &exoplanet.DistanceFromEarth, &exoplanet.Radius, &exoplanet.Mass, &exoplanet.PlanetTypeID, &exoplanet.PlanetType)
	if err != nil && err != sql.ErrNoRows {
		log.Println("GetExoplanetByID err in scanning selectExoplanetByIDQuery : ", err.Error())
		return nil, err
	}

	return &exoplanet, nil

}
