package service

import (
	"errors"
	"log"
	"lumoshive-be-chap38-39/model"
	"lumoshive-be-chap38-39/repository"
	"math"

	"gorm.io/gorm"
)

type ShippingCourierService interface {
	GetAllShippings() ([]model.ShippingCourier, error)
	GetShippingById(id int) (*model.ShippingCourier, error)
	CalculateShippingCost(costReq model.ShippingCostRequest) (*model.ShippingCostResponse, error)
	CalculateCost(distance float64, quantity int) float64
}

type shippingCourierService struct {
	Repo repository.ShippingCourierRepository
}

func NewShippingCourierService(db *gorm.DB) ShippingCourierService {
	return &shippingCourierService{Repo: repository.NewShippingCourierRepository(db)}
}

// GetAllShippings implements ShippingService.
func (s *shippingCourierService) GetAllShippings() ([]model.ShippingCourier, error) {
	return s.Repo.GetAll()
}

// GetShippingById implements ShippingService.
func (s *shippingCourierService) GetShippingById(id int) (*model.ShippingCourier, error) {
	if id == 0 {
		return nil, errors.New("id cannot be 0")
	}
	return s.Repo.GetByID(id)
}

func (s *shippingCourierService) CalculateShippingCost(costReq model.ShippingCostRequest) (*model.ShippingCostResponse, error) {
	log.Printf("Debug: Get calculation for shipping distance: origin: %s - destination: %s", costReq.OriginLatLong, costReq.DestinationLatLong)
	distance, err := s.Repo.CalculateShippingCost(costReq)
	if err != nil {
		return nil, err
	}

	log.Printf("Debug: Get calculation for shipping cost rate")
	shippingData, err := s.GetShippingById(costReq.ShippingID)
	if err != nil {
		return nil, err
	}
	// Convert from meters to kilometers
	distanceInKM := distance / 1000
	shippingCost := math.Round((distanceInKM*shippingData.CostRate)*100) / 100
	log.Printf("Debug: Get calculation for total shipping cost")
	shippingCost += s.CalculateCost(distanceInKM, costReq.Quantity)

	result := model.ShippingCostResponse{
		Distance: distanceInKM,
		Cost:     shippingCost,
	}

	return &result, nil
}

func (s *shippingCourierService) CalculateCost(distance float64, quantity int) float64 {
	var costPerKm int
	if quantity < 2 {
		costPerKm = 2000
	} else {
		costPerKm = 4000
	}

	result := math.Round((distance*float64(costPerKm))*100) / 100
	return result
}
