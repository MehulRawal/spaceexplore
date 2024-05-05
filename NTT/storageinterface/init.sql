CREATE TABLE planet_types (
    id INT AUTO_INCREMENT PRIMARY KEY,
    planet_type VARCHAR(255) UNIQUE,
    description TEXT
);

INSERT INTO planet_types(planet_type, description) VALUES('Gas Giant','composed of only gaseous compounds');
INSERT INTO planet_types(planet_type, description) VALUES('Terrestrial','earth like planets, a bit more rocky and larger than earth');

CREATE TABLE exoplanets (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) UNIQUE,
    description TEXT,
    distance_from_earth BIGINT,
    radius DECIMAL(10,2),
    mass DECIMAL(20,2),
    planet_type_id INT(10)
);