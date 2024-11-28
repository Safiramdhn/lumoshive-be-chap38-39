package repository

import "gorm.io/gorm"

type MainRepository struct {
	ShippingCourierRepository ShippingCourierRepository
	ShippingRepository        ShippingRepository
	ShippingHistoryRepository ShippingHistoryRepository
}

func NewMainRepository(db *gorm.DB) MainRepository {
	return MainRepository{
		ShippingCourierRepository: NewShippingCourierRepository(db),
		ShippingRepository:        NewShippingRepository(db),
		ShippingHistoryRepository: NewShippingHistoryRepository(db),
	}
}
