package service

import "gorm.io/gorm"

type MainService struct {
	ShippingCourierService ShippingCourierService
	ShippingService        ShippingService
	ShippingHistoryService ShippingHistoryService
}

func NewMainService(db *gorm.DB) MainService {
	return MainService{
		ShippingCourierService: NewShippingCourierService(db),
		ShippingService:        NewShippingService(db),
		ShippingHistoryService: NewShippingHistoryService(db),
	}
}
