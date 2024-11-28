package service

import (
	"lumoshive-be-chap38-39/model"
	"lumoshive-be-chap38-39/repository"

	"gorm.io/gorm"
)

type ShippingService interface {
	CreateShipment(shippingInput model.Shipping) (*model.Shipping, error)
}

type shippingService struct {
	Repo repository.ShippingRepository
}

func NewShippingService(db *gorm.DB) ShippingService {
	return &shippingService{Repo: repository.NewShippingRepository(db)}
}

func (s *shippingService) CreateShipment(shippingInput model.Shipping) (*model.Shipping, error) {
	return s.Repo.Create(shippingInput)
}
