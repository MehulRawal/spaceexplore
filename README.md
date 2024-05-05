Space Explore allows explorers to add exoplanets as they discover them and categorize it into particular planettype as they explore.

I've a table for planet_types that helps explorer to store the planet_type and the description, it becomes easier if he introduces more planet types.
SQL : 
  CREATE TABLE planet_types (
      id INT AUTO_INCREMENT PRIMARY KEY,
      planet_type VARCHAR(255) UNIQUE,
      description TEXT
  );


Explorers could store exoplanet data such as name, description, distance from the earth, radius, mass, planet type.
SQL : 
CREATE TABLE exoplanets (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) UNIQUE,
    description TEXT,
    distance_from_earth BIGINT,
    radius DECIMAL(10,2),
    mass DECIMAL(20,2),
    planet_type_id INT(10)
);

The code used gin for routing : go get github.com/gin-gonic/gin

NTT/apis.go contains all the following APIs : 

1. GetPlanetTypes ( returns all the planet_types)
2. AddPlanetTypes ( inputs are planet_type and description,  planet_type must be unique, validation added)
3. GetExoplanets  ( returns all the exoplanets, filtering possible by radius and mass)
4. GetExoplanet ( returns single record for requested id, send id in param)
5. AddExoplanet  ( inputs are name,description, distance from earth, radius , mass, planet_type_id, name must be unique, validations added for name, planet type and if radius == 0 )
6. UpdateExoplanet ( inputs are name,description, distance from earth, radius , mass, planet_type_id, validations same as insert)
7. DeleteExoplanet ( deletes records for given id)
8. GetFuelEstimation ( inputs are exoplanet id and crew capacity, distinguishes which planet type it is and calculates gravity and fuel estimation accordingly)

Fuel estimation to reach an exoplanet can be calculated as :
f = d / (g^2) * c units
d -> distance of exoplanet from earth
g -> gravity of exoplanet
c -> crew capacity (int)

Logic to calculate gravity for each type is as follows :
• Gas Giant :
g = (0.5/r^2)
• Terrestrial :
g = (m/r^2)
m -> mass
r -> radius

NTT/storageinterface contains the models package that contains db connection and db calls.

Response structure for all apis is :
type Response struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}  
to maintain uniformity.



[local spaceexplore.postman_collection.json](https://github.com/MehulRawal/spaceexplore/files/15212932/local.spaceexplore.postman_collection.json)




