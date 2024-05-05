package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func DBInit() gin.HandlerFunc {

	return func(c *gin.Context) {

		// dbHost := os.Getenv("DB_HOST")
		// dbPort := os.Getenv("DB_PORT")
		dbUser := "root"
		dbPass := "NttRawal@123"
		dbName := "SpaceExplore"

		dbHost := "localhost"
		dbPort := "3306"

		// Construct database connection string
		dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

		db, err := sql.Open("mysql", dbURI)
		if err != nil {
			log.Println("InitializeDB() error : ", err.Error())
		}

		c.Set("db", db)
	}
	// Establish a connection to MySQL database

}

func DBClose() gin.HandlerFunc {

	return func(c *gin.Context) {
		db := c.MustGet("db")
		_ = db.(*sql.DB).Close()
	}
}
