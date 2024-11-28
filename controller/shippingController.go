package controller

import (
	"lumoshive-be-chap38-39/model"
	"lumoshive-be-chap38-39/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShippingController struct {
	Service service.ShippingService
}

func NewShippingController(db *gorm.DB) ShippingController {
	return ShippingController{Service: service.NewShippingService(db)}
}

func (c *ShippingController) CreateShippingController(ctx *gin.Context) {
	// get shipping input from body request
	var input model.Shipping
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// create shipping
	shipping, err := c.Service.CreateShipment(input)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// return created shipping
	ctx.JSON(201, gin.H{"data": shipping})
}
