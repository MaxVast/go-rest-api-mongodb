package handler

import (
	"context"
	"encoding/json"
	"github.com/MaxVast/go-rest-api-mongodb/database"
	"github.com/MaxVast/go-rest-api-mongodb/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strconv"
)

type ApiResponse struct {
	Results int                `json:"results"`
	Data    []models.Customers `json:"data"`
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	cursor, err := database.Collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var customers []models.Customers
	if err = cursor.All(ctx, &customers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	results := len(customers)

	response := ApiResponse{
		Results: results,
		Data:    customers,
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idClient := vars["idClient"]

	idClientInt, err := strconv.Atoi(idClient)
	if err != nil {
		http.Error(w, "Invalid idClient format", http.StatusBadRequest)
		return
	}

	filter := bson.M{"id_client": idClientInt}

	ctx := context.Background()

	var customer models.Customer

	err = database.Collection.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Customer not found", http.StatusNotFound)
			return
		}
		log.Printf("Erreur lors de la récupération du client : %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
