package main

import (
	"github.com/gin-gonic/gin"
	"github.com/verssache/go-hacktiv8-2/config"
	"github.com/verssache/go-hacktiv8-2/handler"
	"github.com/verssache/go-hacktiv8-2/helper"
	"github.com/verssache/go-hacktiv8-2/orders"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/verssache/go-hacktiv8-2/docs"
)

// @title Go Hacktiv8 Assignment 2 API Documentation
// @description This is a sample server for a store.
// @termsOfService http://swagger.io/terms/
// @contact.name Gidhan Bagus Algary
// @contact.email gidhanbagusalgary@gmail.com
// @host 127.0.0.1:8080
// @BasePath /api/v1
// @version 1.0.0
// @schemes http
func main() {
	cfg := config.LoadConfig()
	db := helper.InitializeDB()

	orderRepository := orders.NewRepository(db)
	orderService := orders.NewService(orderRepository)
	orderHandler := handler.NewHandler(orderService)

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	api := router.Group("/api/v1")

	api.GET("/orders", orderHandler.FindAll)
	api.GET("/orders/:id", orderHandler.FindByID)
	api.POST("/orders", orderHandler.Save)
	api.PUT("/orders/:id", orderHandler.Update)
	api.DELETE("/orders/:id", orderHandler.Delete)

	// Tugas 1
	api.GET("orders/person/:id", gin.BasicAuth(gin.Accounts{
		cfg.Auth.Username: cfg.Auth.Password,
	}), orderHandler.FindOrderPerson)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := router.Run(":" + cfg.ServerPort)
	if err != nil {
		return
	}
}
