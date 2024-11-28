package controller

import (
	"log"
	"lumoshive-be-chap38-39/model"
	"lumoshive-be-chap38-39/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShippingCourierController struct {
	Service service.MainService
}

func NewShippingCourierController(db *gorm.DB) ShippingCourierController {
	return ShippingCourierController{Service: service.NewMainService(db)}
}

func (c ShippingCourierController) GetAllShippingCourierController(ctx *gin.Context) {
	shippingList, err := c.Service.ShippingCourierService.GetAllShippings()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error getting shipping list": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": shippingList})
}

func (c ShippingCourierController) GetShippingCourierByIdController(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	shippingId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	shipping, err := c.Service.ShippingCourierService.GetShippingById(shippingId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error getting shipping": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": shipping})
}

func (c ShippingCourierController) GetShippingCostController(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	originLatLong := ctx.Param("origin_longlat")
	destinationLatLong := ctx.Param("destination_longlat")
	quantity := ctx.Param("quantity")
	if originLatLong == "" || destinationLatLong == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "origin and destination latitude/longitude are required"})
		return
	}
	shippingId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	ItemQuantity, err := strconv.Atoi(quantity)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid quantity"})
		return
	}

	input := model.ShippingCostRequest{
		ShippingID:         shippingId,
		Quantity:           ItemQuantity,
		OriginLatLong:      originLatLong,
		DestinationLatLong: destinationLatLong,
	}

	log.Printf("Debug: Get calculation for shipping cost")
	shippingCost, err := c.Service.ShippingCourierService.CalculateShippingCost(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error calculating shipping cost": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": shippingCost})
}
