package main

import (
	"github.com/gin-gonic/gin"
	"github.com/verssache/go-hacktiv8-2/config"
	"github.com/verssache/go-hacktiv8-2/handler"
	"github.com/verssache/go-hacktiv8-2/helper"
	"github.com/verssache/go-hacktiv8-2/orders"
)

func main() {
	cfg := config.LoadConfig()
	db := helper.InitializeDB()

	orderRepository := orders.NewRepository(db)
	orderService := orders.NewService(orderRepository)
	orderHandler := handler.NewHandler(orderService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.GET("/orders", orderHandler.FindAll)
	api.GET("/orders/:id", orderHandler.FindByID)
	api.POST("/orders", orderHandler.Save)
	api.PUT("/orders/:id", orderHandler.Update)
	api.DELETE("/orders/:id", orderHandler.Delete)

	router.Run(":" + cfg.ServerPort)

	// Gidhan Bagus Algary
}
