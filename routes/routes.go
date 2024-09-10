package routes

import (
	"github.com/MaxVast/go-rest-api-mongodb/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	r.GET("/customers", handler.GetAllCustomers)
	r.GET("/customer/:idClient", handler.GetCustomerByID)
	r.GET("/funding-types", handler.GetFundingTypes)
}
