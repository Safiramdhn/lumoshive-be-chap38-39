package service

import (
	"lumoshive-be-chap38-39/model"
	"lumoshive-be-chap38-39/repository"

	"gorm.io/gorm"
)

type ShippingService interface {
	CreateShipment(shippingInput model.Shipping) (*model.Shipping, error)
	GetShippingById(id uint) (*model.Shipping, error)
}

type shippingService struct {
	Repo repository.MainRepository
}

func NewShippingService(db *gorm.DB) ShippingService {
	return &shippingService{Repo: repository.NewMainRepository(db)}
}

func (s *shippingService) CreateShipment(shippingInput model.Shipping) (*model.Shipping, error) {
	shippingHistory := model.ShippingHistory{
		Status:   "Created",
		Location: "Waiting for courier to pickup the package",
	}
	shippingInput.ShippingHistory = append(shippingInput.ShippingHistory, shippingHistory)
	shipping, err := s.Repo.ShippingRepository.Create(shippingInput)
	if err != nil {
		return nil, err
	}

	return s.GetShippingById(shipping.ID)
}

func (s *shippingService) GetShippingById(id uint) (*model.Shipping, error) {
	shipping, err := s.Repo.ShippingRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return shipping, nil
}
