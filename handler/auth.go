package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/verssache/go-hacktiv8-2/auth"
	"github.com/verssache/go-hacktiv8-2/helper"
)

type authHandler struct {
	service auth.Service
}

func NewAuthHandler(service auth.Service) *authHandler {
	return &authHandler{service}
}

func (h *authHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
		c.Next()
	}
}
