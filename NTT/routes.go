package main

import (
	api "NTT/apis.go"

	models "NTT/storageinterface"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() *gin.Engine {
	router := gin.Default()

	// Define route handlers
	router.Use(models.DBInit())
	router.GET("/planet_types", api.GetPlanetTypes)
	router.PUT("/planet_types", api.AddPlanetTypes)
	router.GET("/exoplanets", api.GetExoplanets)
	router.GET("/exoplanet/:id", api.GetExoplanet)
	router.PUT("/exoplanet", api.AddExoplanet)
	router.PATCH("/exoplanet/:id", api.UpdateExoplanet)
	router.DELETE("/exoplanet/:id", api.DeleteExoplanet)
	router.POST("/exoplanet/fuel_estimate", api.GetFuelEstimation)
	router.Use(models.DBClose())
	return router
}
