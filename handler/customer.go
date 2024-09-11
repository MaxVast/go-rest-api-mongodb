package handler

import (
	"context"
	"errors"
	"github.com/MaxVast/go-rest-api-mongodb/database"
	"github.com/MaxVast/go-rest-api-mongodb/models"
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
