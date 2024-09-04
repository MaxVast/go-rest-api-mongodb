package handler

import (
	"context"
	"encoding/json"
	"github.com/MaxVast/go-rest-api-mongodb/database"
	"github.com/MaxVast/go-rest-api-mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type ApiResponse struct {
	Results int               `json:"results"`
	MaxPage int               `json:"max_page"`
	Data    []models.Customer `json:"data"`
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	cursor, err := database.Collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var customers []models.Customer
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
