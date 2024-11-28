package service

import (
	"errors"
	"lumoshive-be-chap38-39/model"
	"lumoshive-be-chap38-39/repository"

	"gorm.io/gorm"
)

type ShippingHistoryService interface {
	CreateShippingHistory(historyInput model.ShippingHistory) error
	GetShippingHistoryByShippingID(shippingID uint) ([]model.ShippingHistory, error)
}

type shippingHistoryService struct {
	Repo repository.MainRepository
}

func NewShippingHistoryService(db *gorm.DB) ShippingHistoryService {
	return &shippingHistoryService{Repo: repository.NewMainRepository(db)}
}

func (s *shippingHistoryService) CreateShippingHistory(historyInput model.ShippingHistory) error {
	if historyInput.ShippingID == 0 || historyInput.Status == "" {
		return errors.New("missing required fields: shipping ID and status are mandatory")
	}
	return s.Repo.ShippingHistoryRepository.Create(historyInput)
}

func (s *shippingHistoryService) GetShippingHistoryByShippingID(shippingID uint) ([]model.ShippingHistory, error) {
	return s.Repo.ShippingHistoryRepository.GetByShippingID(shippingID)
}
