package main

import (
	"fetch-rewards/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/receipts/process", controller.PostReceipts)
	router.GET("/receipts/:id/points", controller.GetPoints)
	router.Run("localhost:8080")
}
