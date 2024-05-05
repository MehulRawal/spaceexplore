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
	router.Use(models.DBClose())
	return router
}
