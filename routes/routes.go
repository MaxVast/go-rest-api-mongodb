package routes

import (
	"github.com/MaxVast/go-rest-api-mongodb/handler"
	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/customers", handler.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customer/{idClient}", handler.GetCustomerByID).Methods("GET")
	return router
}
