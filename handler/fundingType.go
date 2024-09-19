package handler

import (
	"fmt"
	"github.com/MaxVast/go-rest-api-mongodb/database"
	"github.com/MaxVast/go-rest-api-mongodb/models/sqlServer"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetFundingTypes(c *gin.Context) {
	// Prepare PS
	rows, err := database.DB.Query("EXEC ps_API_L_TYPE_FINANCEMENT")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error while executing the procedure : %v", err)})
		return
	}
	defer rows.Close()

	var fundingTypes []sqlServer.FundingType
	for rows.Next() {
		var ft sqlServer.FundingType
		if err := rows.Scan(&ft.ID, &ft.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error reading results : %v", err)})
			return
		}
		fundingTypes = append(fundingTypes, ft)
	}

	totalItems := len(fundingTypes)

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error processing results : %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"totalItems": totalItems, "funding_types": fundingTypes})
}
