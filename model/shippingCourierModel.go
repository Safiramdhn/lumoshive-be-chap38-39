package model

import (
	"time"

	"gorm.io/gorm"
)

type ShippingCourier struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name"`
	CostRate  float64        `json:"rate_cost" gorm:"type:decimal(10,2);not null"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type ShippingCostRequest struct {
	ShippingID         int    `json:"id_shipping"`
	Quantity           int    `json:"quantity_barang"`
	OriginLatLong      string `json:"origin_latlong"`
	DestinationLatLong string `json:"destination_latlong"`
}

type ShippingCostResponse struct {
	Distance float64 `json:"distance"`
	Cost     float64 `json:"cost"`
}
