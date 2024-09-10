package main

import (
	"github.com/MaxVast/go-rest-api-mongodb/database"
	"github.com/MaxVast/go-rest-api-mongodb/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func main() {
	// Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file : %v", err)
	}

	// Init connection to BDD
	database.Connect()

	database.ConnectSql()

	// Init Gin router
	r := gin.Default()

	// Define the trusted Proxies (example : 127.0.0.1, localhost)
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Init CORS settings
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Init routes API
	routes.InitializeRoutes(r)

	// Start the server
	log.Fatal(r.Run("127.0.0.1:8000"))
}
