package database

import (
	"fmt"
	"log"
	"lumoshive-be-chap38-39/model"

	"gorm.io/gorm"
)

func initiateShippingData(db *gorm.DB) error {
	// Check if the data already exists
	var count int64
	if err := db.Model(&model.ShippingCourier{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to count existing shipping data: %w", err)
	}

	if count > 0 {
		log.Println("Shipping data already initialized, skipping seeder.")
		return nil
	}

	// Predefined shipping courier data
	shippingData := []model.ShippingCourier{
		{Name: "JNE", CostRate: 5.00},
		{Name: "JNT", CostRate: 10.00},
		{Name: "Pos Indonesia", CostRate: 15.00},
	}

	// Insert data into the database
	if err := db.Create(&shippingData).Error; err != nil {
		return fmt.Errorf("failed to initialize shipping data: %w", err)
	}

	log.Println("Shipping data initialized successfully.")
	return nil
}
