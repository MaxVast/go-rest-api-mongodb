package main

import (
	"github.com/MaxVast/go-rest-api-mongodb/database"
	"github.com/MaxVast/go-rest-api-mongodb/routes"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	// Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env : %v", err)
	}

	// Init connection to BDD
	database.Connect()

	// Init routes API
	router := routes.InitializeRoutes()
	log.Fatal(http.ListenAndServe(":8000", router))
}
