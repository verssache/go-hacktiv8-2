package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/verssache/go-hacktiv8-2/helper"
	"github.com/verssache/go-hacktiv8-2/orders"
)

type orderHandler struct {
	service orders.Service
}

func NewHandler(service orders.Service) *orderHandler {
	return &orderHandler{service}
}

// @BasePath /api/v1

// FindAll godoc
// @Summary Show all orders
// @Description Get all orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {json} helper.APIResponse
// @Router /orders [get]
func (h *orderHandler) FindAll(c *gin.Context) {
	findOrders, err := h.service.FindAll()
	if err != nil {
		response := helper.APIResponse("Error to get orders", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of orders", http.StatusOK, "success", orders.FormatOrders(findOrders))
	c.JSON(http.StatusOK, response)
}

// FindAllByUserID godoc
// @Summary Show all orders by user id
// @Description Get all orders by user id
// @Tags orders
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID"
// @Success 200 {json} helper.APIResponse
// @Router /orders/user/{user_id} [get]
func (h *orderHandler) FindByID(c *gin.Context) {
	var input orders.FindOrderInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to get detail of order", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	findOrder, err := h.service.FindByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Detail of order", http.StatusOK, "success", orders.FormatOrder(findOrder))
	c.JSON(http.StatusOK, response)
}

// Save godoc
// @Summary Create new order
// @Description Create new order
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body orders.SaveOrderInput true "Order"
// @Success 200 {json} helper.APIResponse
// @Router /orders [post]
func (h *orderHandler) Save(c *gin.Context) {
	var input orders.SaveOrderInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to create order", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newOrder, err := h.service.Save(input)
	if err != nil {
		response := helper.APIResponse("Failed to create order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create order", http.StatusOK, "success", orders.FormatOrder(newOrder))
	c.JSON(http.StatusOK, response)
}

// Update godoc
// @Summary Update order
// @Description Update order
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order_id path int true "Order ID"
// @Param order body orders.UpdateOrderInput true "Order"
// @Success 200 {json} helper.APIResponse
// @Router /orders/{order_id} [put]
func (h *orderHandler) Update(c *gin.Context) {
	var inputID orders.FindOrderInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to update order", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData orders.UpdateOrderInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to update order", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedOrder, err := h.service.Update(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update order", http.StatusOK, "success", orders.FormatOrder(updatedOrder))
	c.JSON(http.StatusOK, response)
}

// Delete godoc
// @Summary Delete order
// @Description Delete order
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order_id path int true "Order ID"
// @Success 200 {json} helper.APIResponse
// @Router /orders/{order_id} [delete]
func (h *orderHandler) Delete(c *gin.Context) {
	var input orders.FindOrderInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to delete order", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.Delete(input)
	if err != nil {
		response := helper.APIResponse("Failed to delete order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete order", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
