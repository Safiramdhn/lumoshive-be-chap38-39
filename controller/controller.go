package controller

import (
	"gorm.io/gorm"
)

type MainController struct {
	ShippingCourierController ShippingCourierController
	ShippingController        ShippingController
}

func NewMainController(db *gorm.DB) MainController {
	return MainController{
		ShippingCourierController: NewShippingCourierController(db),
		ShippingController:        NewShippingController(db),
	}
}
