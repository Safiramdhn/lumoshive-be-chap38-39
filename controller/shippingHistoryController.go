package controller

import (
	"lumoshive-be-chap38-39/model"
	"lumoshive-be-chap38-39/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShippingHistoryController struct {
	Service service.MainService
}

func NewShippingHistoryController(db *gorm.DB) ShippingHistoryController {
	return ShippingHistoryController{Service: service.NewMainService(db)}
}

func (c *ShippingHistoryController) UpdateShippingHistoryController(ctx *gin.Context) {
	idParam := ctx.Param("shipping_id")

	// Validate and parse the shipping ID
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid shipping ID. ID must be a positive integer.",
		})
		return
	}

	var newHistory model.ShippingHistory
	// Bind and validate input data
	if err := ctx.ShouldBindJSON(&newHistory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	// Set the shipping ID for the new history entry
	newHistory.ShippingID = uint(id)

	// Attempt to create the shipping history
	if err := c.Service.ShippingHistoryService.CreateShippingHistory(newHistory); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update shipping history: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Shipping history has been updated successfully.",
	})
}

func (c *ShippingHistoryController) GetShippingHistoryController(ctx *gin.Context) {
	shippingIDParam := ctx.Param("shipping_id")
	id, err := strconv.Atoi(shippingIDParam)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid shipping ID. ID must be a positive integer.",
		})
	}

	shiipingID := uint(id)
	shippingHistory, err := c.Service.ShippingHistoryService.GetShippingHistoryByShippingID(shiipingID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Shipping history not found."})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Shipping history successfully retrieved",
		"data":    shippingHistory})
}
