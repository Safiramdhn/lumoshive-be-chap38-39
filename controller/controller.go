package controller

import (
	"gorm.io/gorm"
)

type MainController struct {
	ShippingCourierController ShippingCourierController
	ShippingController        ShippingController
	ShippingHistoryController ShippingHistoryController
}

func NewMainController(db *gorm.DB) MainController {
	return MainController{
		ShippingCourierController: NewShippingCourierController(db),
		ShippingController:        NewShippingController(db),
		ShippingHistoryController: NewShippingHistoryController(db),
	}
}
