package database

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"os"
)

var DB *sql.DB

func ConnectSql() {
	sqlsvr := os.Getenv("DATABASE_URL")

	if sqlsvr == "" {
		log.Fatal("DATABASE_URL not defined in .env file")
	}

	var err error
	DB, err = sql.Open("sqlserver", sqlsvr)
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
	}

	// Check connexion
	if err = DB.Ping(); err != nil {
		log.Fatalf("Error checking connection to SQL Server: %v", err)
	}

	log.Println("Connected to SQL Server.")
}
