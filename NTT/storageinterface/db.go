package models

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

func DBInit() gin.HandlerFunc {

	return func(c *gin.Context) {

		db, err := sql.Open("mysql", "test_user:secret@tcp(mysql:3306)/test_database")
		if err != nil {
			log.Println("error in selectPlanetTypeQuery : ", err.Error())
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
