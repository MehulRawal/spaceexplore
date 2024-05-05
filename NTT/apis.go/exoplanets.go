package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	models "NTT/storageinterface"

	"github.com/gin-gonic/gin"
)

type ExoplanetFilterRequest struct {
	Filters *models.ExoplanetFilters `json:"filters,omitempty"`
}

func GetPlanetTypes(c *gin.Context) {

	db := c.MustGet("db").(*sql.DB)
	planetTypes, err := models.GetPlanetTypes(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, planetTypes)
}

func AddPlanetTypes(c *gin.Context) {

	db := c.MustGet("db").(*sql.DB)

	var req models.PlanetType

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.PlanetType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "", "message": "Please provide valid planet_type"})
		return
	}

	isPlanetTypeValid, err := models.CheckIfPlanetTypeExist(db, req.PlanetType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if isPlanetTypeValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "", "message": fmt.Sprintf("Planet Type %s already exist", req.PlanetType)})
		return
	}

	planetTypeID, err := models.InsertPlanetTypes(db, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, planetTypeID)
}

func GetExoplanets(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	req := ExoplanetFilterRequest{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println("GetExoplanets API error in ShouldBindJSON : ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	exoPlanetlist, err := models.GetExoplanets(db, req.Filters)
	if err != nil {
		log.Println("GetExoplanets API error in GetExoplanets : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	c.JSON(http.StatusOK, exoPlanetlist)

}

func GetExoplanet(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	idStr := c.Param("id")

	var exoplanetId int

	var err error

	if idStr != "" {
		exoplanetId, err = strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "", "message": fmt.Sprintf("Invalid id : %d", idStr)})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "", "message": fmt.Sprintf("Invalid id : %d", idStr)})
		return
	}

	if exoplanetId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "", "message": "id cannot be zero"})
		return
	}

	exoPlanetlist, err := models.GetExoplanetByID(db, exoplanetId)
	if err != nil {
		log.Println("GetExoplanet API error in GetExoplanetByID : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	c.JSON(http.StatusOK, exoPlanetlist)

}

func AddExoplanet(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var req models.ExoPlanet

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "", "message": "Please provide valid name for the exoplanet"})
		return
	}

	isNameValid, err := models.CheckIfNameAlreadyExist(db, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if isNameValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "", "message": fmt.Sprintf("Planet with name %s already exist", req.Name)})
		return
	}

	if req.Radius == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "", "message": "Planet cannot have radii zero"})
		return
	}

	if req.PlanetTypeID != 0 {
		isPlanetTypeValid, err := models.CheckIfPlanetTypeByIDExist(db, req.PlanetTypeID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if !isPlanetTypeValid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "", "message": fmt.Sprintf("Planet Type %s does not exist, Kindly Add %s to the Planet Type list", req.PlanetType, req.PlanetType)})
			return
		}
	}

	exoplanetID, err := models.InsertExoplanet(db, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exoplanetID)
}

func UpdateExoplanet(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	idStr := c.Param("id")

	var exoplanetId int

	var err error

	if idStr != "" {
		exoplanetId, err = strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "", "message": fmt.Sprintf("Invalid id : %d", idStr)})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "", "message": fmt.Sprintf("Invalid id : %d", idStr)})
		return
	}

	var req models.ExoPlanet

	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if exoplanetId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "", "message": "id cannot be zero"})
		return
	}

	if exoplanetId != 0 {
		exoplanetExists, err := models.CheckIfExoplanetExist(db, exoplanetId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if !exoplanetExists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "", "message": fmt.Sprintf("Exoplanet does not exist for given id : %d", req.ID)})
			return
		}
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "", "message": "Please provide valid name for the exoplanet"})
		return
	}

	isNameValid, err := models.CheckIfNameAlreadyExist(db, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if isNameValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "", "message": fmt.Sprintf("Planet with name %s already exist", req.Name)})
		return
	}

	err = models.UpdateExoplanet(db, exoplanetId, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
