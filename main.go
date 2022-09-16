package main

import (
	"github.com/gin-gonic/gin"
	"github.com/verssache/go-hacktiv8-2/auth"
	"github.com/verssache/go-hacktiv8-2/config"
	"github.com/verssache/go-hacktiv8-2/handler"
	"github.com/verssache/go-hacktiv8-2/helper"
	"github.com/verssache/go-hacktiv8-2/orders"
	"github.com/verssache/go-hacktiv8-2/users"

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

	authService := auth.NewService()

	orderRepository := orders.NewRepository(db, cfg.Host)
	userRepository := users.NewRepository(db)
	orderService := orders.NewService(orderRepository)
	userService := users.NewService(userRepository)
	orderHandler := handler.NewHandler(orderService)
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	api := router.Group("/api/v1")

	api.GET("/orders", orderHandler.FindAll)
	api.GET("/orders/:id", orderHandler.FindByID)
	api.POST("/orders", authService.AuthMiddleware(authService, userService), orderHandler.Save)
	api.PUT("/orders/:id", authService.AuthMiddleware(authService, userService), orderHandler.Update)
	api.DELETE("/orders/:id", authService.AuthMiddleware(authService, userService), orderHandler.Delete)

	api.GET("orders/person/:id", gin.BasicAuth(gin.Accounts{
		cfg.Auth.Username: cfg.Auth.Password,
	}), orderHandler.FindOrderPerson)

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := router.Run(":" + cfg.ServerPort)
	if err != nil {
		return
	}
}
