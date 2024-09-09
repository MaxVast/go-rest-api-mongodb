package main

import (
	"github.com/MaxVast/go-rest-api-mongodb/database"
	"github.com/MaxVast/go-rest-api-mongodb/routes"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	// Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env : %v", err)
	}

	//Init CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:3000/"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Accept"},
		AllowCredentials: true,
	})

	// Init connection to BDD
	database.Connect()

	// Init routes API
	router := routes.InitializeRoutes()
	//Apply CORS to ROUTER
	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe("127.0.0.1:8000", handler))
}
