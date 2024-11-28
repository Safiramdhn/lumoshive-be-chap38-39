package model

import (
	"time"

	"gorm.io/gorm"
)

type Shipping struct {
	ID                 uint              `json:"id" gorm:"primaryKey;autoIncrement"`
	TransactionID      string            `json:"transaction_id" gorm:"type:varchar(255);not null"`            // Define the type and make it not null
	OriginLatLong      string            `json:"origin_latlong" gorm:"type:varchar(50);not null"`             // Define the type and size
	DestinationLatLong string            `json:"destination_latlong" gorm:"type:varchar(50);not null"`        // Define the type and size
	TotalShippingCost  float64           `json:"total_shipping_cost" gorm:"type:decimal(10,2);not null"`      // Decimal with precision
	ShippingHistory    []ShippingHistory `json:"shipping_history" gorm:"foreignKey:ShippingID;references:ID"` // Define a one-to-many relationship with ShippingHistory
	CreatedAt          time.Time         `json:"-" gorm:"autoCreateTime"`                                     // Automatically set the creation time
	UpdatedAt          time.Time         `json:"-" gorm:"autoUpdateTime"`                                     // Automatically set the update time
	DeletedAt          gorm.DeletedAt    `json:"-" gorm:"index"`                                              // Soft delete with an index
}
