package main

import (

	// "encoding/json"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := SetUpRoutes()

	// Start the HTTP server
	log.Println("Server listening on :8080")
	log.Fatal(router.Run(":8080"))
}
