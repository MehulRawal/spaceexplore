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

type FuelEstimationRequest struct {
	ToExoplanetID int `json:"to_exoplanet_id"`
	CrewCapacity  int `json:"crew_capacity"`
}

func GetPlanetTypes(c *gin.Context) {

	db := c.MustGet("db").(*sql.DB)
	planetTypes, err := models.GetPlanetTypes(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseConstruct("", err.Error(), false, nil))
		return
	}

	c.JSON(http.StatusOK, ResponseConstruct("", "", true, planetTypes))
}

func AddPlanetTypes(c *gin.Context) {

	db := c.MustGet("db").(*sql.DB)

	var req models.PlanetType

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseConstruct("", err.Error(), false, nil))
		return
	}

	if req.PlanetType == "" {
		c.JSON(http.StatusBadRequest, ResponseConstruct("Please provide valid planet_type", "", false, nil))
	}

	isPlanetTypeValid, err := models.CheckIfPlanetTypeExist(db, req.PlanetType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseConstruct("", err.Error(), false, nil))
		return
	}

	if isPlanetTypeValid {
		c.JSON(http.StatusBadRequest, ResponseConstruct(fmt.Sprintf("Planet Type %s already exist", req.PlanetType), "", false, nil))
		return
	}

	planetTypeID, err := models.InsertPlanetTypes(db, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseConstruct("", err.Error(), false, nil))
		return
	}

	c.JSON(http.StatusOK, ResponseConstruct("Successfully added planet type", "", true, planetTypeID))
}

func GetExoplanets(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	req := ExoplanetFilterRequest{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println("GetExoplanets API error in ShouldBindJSON : ", err.Error())
		c.JSON(http.StatusBadRequest, ResponseConstruct("Invalid Request", err.Error(), false, nil))
		return
	}

	exoPlanetlist, err := models.GetExoplanets(db, req.Filters)
	if err != nil {
		log.Println("GetExoplanets API error in GetExoplanets : ", err.Error())
		c.JSON(http.StatusInternalServerError, ResponseConstruct("", err.Error(), false, nil))
		return
	}

	c.JSON(http.StatusOK, ResponseConstruct("", "", true, exoPlanetlist))

}

func GetExoplanet(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	idStr := c.Param("id")

	var exoplanetId int

	var err error

	if idStr != "" {
		exoplanetId, err = strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, ResponseConstruct(fmt.Sprintf("Invalid id : %s", idStr), err.Error(), false, nil))
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, ResponseConstruct("id cannot be empty", "", false, nil))
		return
	}

	if exoplanetId == 0 {
		c.JSON(http.StatusBadRequest, ResponseConstruct("id cannot be zero", "", false, nil))
		return
	}

	exoPlanetlist, err := models.GetExoplanetByID(db, exoplanetId)
	if err != nil {
		log.Println("GetExoplanet API error in GetExoplanetByID : ", err.Error())
		c.JSON(http.StatusInternalServerError, ResponseConstruct("", err.Error(), false, nil))
		return
	}

	c.JSON(http.StatusOK, ResponseConstruct("", "", true, exoPlanetlist))

}

func AddExoplanet(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var req models.ExoPlanet

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseConstruct("", err.Error(), false, nil))
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, ResponseConstruct("Please provide valid name for the exoplanet", "", false, nil))
		return
	}

	isNameValid, err := models.CheckIfNameAlreadyExist(db, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseConstruct("", err.Error(), false, nil))
		return
	}

	if isNameValid {
		c.JSON(http.StatusBadRequest, ResponseConstruct(fmt.Sprintf("Planet with name %s already exist", req.Name), "", false, nil))
		return
	}

	if req.Radius == 0 {
		c.JSON(http.StatusBadRequest, ResponseConstruct("Planet cannot have radii zero", "", false, nil))
		return
	}

	if req.PlanetTypeID != 0 {
		isPlanetTypeValid, err := models.CheckIfPlanetTypeByIDExist(db, req.PlanetTypeID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ResponseConstruct("", err.Error(), false, nil))
			return
		}

		if !isPlanetTypeValid {
			c.JSON(http.StatusBadRequest, ResponseConstruct(fmt.Sprintf("Planet Type %s does not exist, Kindly Add %s to the Planet Type list", req.PlanetType, req.PlanetType), "", false, nil))
			return
		}

		// mass is zero for Gas Giants
		if req.PlanetTypeID == 1 {
			req.Mass = 0
		}
	}

	exoplanetID, err := models.InsertExoplanet(db, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseConstruct("", err.Error(), false, nil))
		return
	}

	c.JSON(http.StatusOK, ResponseConstruct("Successfully Added Exoplanet", "", true, exoplanetID))
}

func UpdateExoplanet(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	idStr := c.Param("id")

	var exoplanetId int

	var err error

	if idStr != "" {
		exoplanetId, err = strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, ResponseConstruct(fmt.Sprintf("Invalid id : %s", idStr), err.Error(), false, nil))
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, ResponseConstruct(fmt.Sprintf("Invalid id : %s", idStr), "", false, nil))
		return
	}

	var req models.ExoPlanet

	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseConstruct("", err.Error(), false, nil))
		return
	}

	if exoplanetId == 0 {
		c.JSON(http.StatusBadRequest, ResponseConstruct("id cannot be zero", "", false, nil))
		return
	}

	if exoplanetId != 0 {
		exoplanetExists, err := models.CheckIfExoplanetExist(db, exoplanetId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ResponseConstruct("", err.Error(), false, nil))
			return
		}

		if !exoplanetExists {
			c.JSON(http.StatusBadRequest, ResponseConstruct(fmt.Sprintf("Exoplanet does not exist for given id : %d", req.ID), "", false, nil))
			return
		}
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, ResponseConstruct("Please provide valid name for the exoplanet", "", false, nil))
		return
	}

	isNameValid, err := models.CheckIfNameAlreadyExistWithOtherId(db, req.Name, exoplanetId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseConstruct("", err.Error(), false, nil))
		return
	}

	if isNameValid {
		c.JSON(http.StatusBadRequest, ResponseConstruct(fmt.Sprintf("Planet with name %s already exist", req.Name), "", false, nil))
		return
	}

	if req.PlanetTypeID != 0 {
		isPlanetTypeValid, err := models.CheckIfPlanetTypeByIDExist(db, req.PlanetTypeID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ResponseConstruct("", err.Error(), false, nil))
			return
		}

		if !isPlanetTypeValid {
			c.JSON(http.StatusBadRequest, ResponseConstruct(fmt.Sprintf("Planet Type %s does not exist, Kindly Add %s to the Planet Type list", req.PlanetType, req.PlanetType), "", false, nil))
			return
		}

		// mass is zero for Gas Giants
		if req.PlanetTypeID == 1 {
			req.Mass = 0
		}
	}

	err = models.UpdateExoplanet(db, exoplanetId, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseConstruct("", err.Error(), false, nil))
		return
	}

	c.JSON(http.StatusOK, ResponseConstruct("Successfully Update Exoplanet", "", true, exoplanetId))
}

func DeleteExoplanet(c *gin.Context) {

	db := c.MustGet("db").(*sql.DB)
	idStr := c.Param("id")

	var exoplanetId int

	var err error

	if idStr != "" {
		exoplanetId, err = strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, ResponseConstruct(fmt.Sprintf("Invalid id : %s", idStr), err.Error(), false, nil))
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, ResponseConstruct(fmt.Sprintf("Invalid id : %s", idStr), "", false, nil))
		return
	}

	err = models.DeleteExoplanetByID(db, exoplanetId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseConstruct("", err.Error(), false, nil))
		return
	}

	c.JSON(http.StatusOK, ResponseConstruct("Successfully Deleted Exoplanet", "", true, nil))
}

func GetFuelEstimation(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var req FuelEstimationRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseConstruct("", err.Error(), false, nil))
		return
	}

	if req.ToExoplanetID == 0 {
		c.JSON(http.StatusBadRequest, ResponseConstruct("Please provide valid exoplanet to travel", "", false, nil))
		return
	}

	if req.ToExoplanetID != 0 {
		exoplanetExists, err := models.CheckIfExoplanetExist(db, req.ToExoplanetID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ResponseConstruct("", err.Error(), false, nil))
			return
		}

		if !exoplanetExists {
			c.JSON(http.StatusBadRequest, ResponseConstruct(fmt.Sprintf("Exoplanet : %d does not exist you wish to travel ", req.ToExoplanetID), "", false, nil))
			return
		}
	}

	fuelEstimate, err := models.FuelEstimation(db, req.ToExoplanetID, req.CrewCapacity)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseConstruct("", err.Error(), false, nil))
		return
	}

	c.JSON(http.StatusOK, ResponseConstruct("", "", true, fmt.Sprintf("Estimated fuel requirement is %.2f for crew capacity : %d", fuelEstimate, req.CrewCapacity)))
}
