package controller

import (
	"fetch-rewards/controller/utils"
	"fetch-rewards/data"
	"fetch-rewards/models"
	"fmt"
	"net/http"

	"fetch-rewards/models/modelutils"

	"github.com/gin-gonic/gin"
)

// POST
// PostReceipts handles the POST request for adding the new receipts
func PostReceipts(c *gin.Context) {
	var newReceipt models.Receipt

	//if JSON does not binds with the Receipt struct, returns the error msg
	if err := c.BindJSON(&newReceipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "The receipt is invalid"})
		return
	}

	//checks for the requires elements in the request body and compares regex pattern
	if err := modelutils.ValidateReceipt(newReceipt); err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	//Generates a unique id associated with the receipt
	id := utils.GenerateUniqueID(newReceipt)

	// Adding the new receipt record to the in-memory data store.
	newRecord := models.Record{ID: id, Receipt: newReceipt}
	data.Records = append(data.Records, newRecord)
	idResponse := models.IDResponse{ID: newRecord.ID}
	c.IndentedJSON(http.StatusCreated, idResponse)
}

// GET by ID request
// Calculation of points takes place at this stage
func GetPoints(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of records, looking for
	// an result whose ID value matches the parameter.
	for _, r := range data.Records {
		if r.ID == id {
			points := utils.CalculatePoints(r.Receipt)
			r.Points = points
			pointsResponse := models.PointsResponse{Points: r.Points}
			c.IndentedJSON(http.StatusOK, pointsResponse)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No receipt found for that id"})
}

// Home Page
func HomePage(c *gin.Context) {
	c.String(http.StatusOK, "Hey there! This is Golang web service project for computing receipt rewards")
}
