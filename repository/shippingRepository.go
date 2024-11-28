package repository

import (
	"lumoshive-be-chap38-39/model"

	"gorm.io/gorm"
)

type ShippingRepository interface {
	Create(shippingInput model.Shipping) (model.Shipping, error)
	GetByID(id uint) (*model.Shipping, error)
}

type shippingRepository struct {
	DB *gorm.DB
}

func NewShippingRepository(db *gorm.DB) ShippingRepository {
	return &shippingRepository{DB: db}
}

func (repo *shippingRepository) Create(shippingInput model.Shipping) (model.Shipping, error) {
	// Save the shipping to the database
	err := repo.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&shippingInput).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return model.Shipping{}, err
	}
	return shippingInput, nil
}

func (repo *shippingRepository) GetByID(id uint) (*model.Shipping, error) {
	// Retrieve the shipping from the database by ID
	var shipping model.Shipping
	err := repo.DB.Preload("ShippingHistory").Where("id = ?", id).Find(&shipping).Error
	if err != nil {
		return nil, err
	}
	return &shipping, nil
}
