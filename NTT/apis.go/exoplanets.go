package api

import (
	"database/sql"
	"net/http"

	models "NTT/storageinterface"

	"github.com/gin-gonic/gin"
)

func GetPlanetTypes(c *gin.Context) {

	db := c.MustGet("db").(*sql.DB)
	planetTypes, err := models.GetPlanetTypes(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, planetTypes)
}
