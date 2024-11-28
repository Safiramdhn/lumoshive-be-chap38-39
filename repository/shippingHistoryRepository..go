package repository

import (
	"fmt"
	"lumoshive-be-chap38-39/model"

	"gorm.io/gorm"
)

type ShippingHistoryRepository interface {
	Create(shippingHistoryInput model.ShippingHistory) error
	GetByShippingID(shippingID uint) ([]model.ShippingHistory, error)
}

type shippingHistoryRepository struct {
	DB *gorm.DB
}

func NewShippingHistoryRepository(db *gorm.DB) ShippingHistoryRepository {
	return &shippingHistoryRepository{DB: db}
}

func (repo *shippingHistoryRepository) Create(shippingHistoryInput model.ShippingHistory) error {
	return repo.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&shippingHistoryInput).Error; err != nil {
			return fmt.Errorf("failed to save shipping history: %w", err)
		}
		return nil
	})
}

func (repo *shippingHistoryRepository) GetByShippingID(shippingID uint) ([]model.ShippingHistory, error) {
	var result []model.ShippingHistory
	err := repo.DB.Preload("Shipping").Where("shipping_id = ?", shippingID).Find(&result).Error
	return result, err
}
