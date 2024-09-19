package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/MaxVast/go-rest-api-mongodb/database"
	"github.com/MaxVast/go-rest-api-mongodb/models"
	"github.com/MaxVast/go-rest-api-mongodb/models/sqlServer"
	"github.com/MaxVast/go-rest-api-mongodb/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
	"net/http"
	"strconv"
)

type ApiResponse struct {
	TotalItems int                `json:"totalItems"`
	Data       []models.Customers `json:"data"`
}

func GetAllCustomers(c *gin.Context) {
	ctx := context.Background()

	cursor, err := database.Collection.Find(ctx, bson.M{})
	if errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	var customers []models.Customers
	if err = cursor.All(ctx, &customers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	results := len(customers)

	response := ApiResponse{
		TotalItems: results,
		Data:       customers,
	}

	// Return DATA JSON
	c.JSON(http.StatusOK, response)
}

func GetCustomerByID(c *gin.Context) {
	idClient := c.Param("idClient")

	idClientInt, err := strconv.Atoi(idClient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid idClient format"})
		return
	}

	filter := bson.M{"id_client": idClientInt}

	ctx := context.Background()

	var customer models.Customer

	err = database.Collection.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
			return
		}
		log.Printf("Erreur lors de la récupération du client : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Return DATA JSON
	c.JSON(http.StatusOK, customer)
}

func GetInfoCustomerByName(c *gin.Context) {
	nameClient := c.Param("name")

	// Prepare PS
	query := "EXEC ps_API_L_CLIENT @nom = @nameParam"
	rows, err := database.DB.Query(query, sql.Named("nameParam", nameClient))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error while executing the procedure : %v", err)})
		return
	}
	defer rows.Close()

	var infoCustomers []sqlServer.InfoCustomers

	for rows.Next() {
		var infoCustomer sqlServer.InfoCustomers
		if err := service.ScanStruct(&infoCustomer, rows); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error scanning data : %v", err)})
			return
		}

		infoCustomers = append(infoCustomers, infoCustomer)
	}

	totalItems := len(infoCustomers)

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error processing results : %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": infoCustomers, "totalItems": totalItems})
}
