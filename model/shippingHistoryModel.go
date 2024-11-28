package model

import (
	"time"

	"gorm.io/gorm"
)

type ShippingHistory struct {
	ID         uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	ShippingID uint           `gorm:"not null" json:"-"`                                              // Ensures the foreign key is not null and indexed
	Shipping   Shipping       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"shipping"` // Ensures the foreign key is not null and indexed
	Status     string         `json:"status" gorm:"type:varchar(255);not null"`                       // Adds a max length and ensures it's not null
	Location   string         `json:"location" gorm:"type:varchar(255);not null"`                     // Adds a max length and ensures it's not null
	CreatedAt  time.Time      `json:"-" gorm:"autoCreateTime"`                                        // Automatically set the creation time
	UpdatedAt  time.Time      `json:"-" gorm:"autoUpdateTime"`                                        // Automatically set the update time
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`                                                 // Soft delete with an index
}
