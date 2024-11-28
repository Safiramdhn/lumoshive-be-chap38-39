package controller

import (
	"lumoshive-be-chap38-39/model"
	"lumoshive-be-chap38-39/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShippingController struct {
	Service service.MainService
}

func NewShippingController(db *gorm.DB) ShippingController {
	return ShippingController{Service: service.NewMainService(db)}
}

func (c *ShippingController) CreateShippingController(ctx *gin.Context) {
	// get shipping input from body request
	var input model.Shipping
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// create shipping
	shipping, err := c.Service.ShippingService.CreateShipment(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// return created shipping
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Shipping created successfully",
		"data":    shipping,
	})
}

func (c *ShippingController) GetShippingByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	// Validate and parse the ID
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid shipping ID. ID must be a positive integer.",
		})
		return
	}

	shippingID := uint(id)

	// Retrieve shipping data
	shipping, err := c.Service.ShippingService.GetShippingById(shippingID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Shipping record not found: " + err.Error(),
		})
		return
	}

	// Respond with the shipping data
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Shipping record found",
		"data":    shipping,
	})
}
