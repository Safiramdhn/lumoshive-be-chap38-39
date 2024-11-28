package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"lumoshive-be-chap38-39/model"
	"net/http"

	"gorm.io/gorm"
)

type ShippingCourierRepository interface {
	GetAll() ([]model.ShippingCourier, error)
	GetByID(id int) (*model.ShippingCourier, error)
	CalculateShippingCost(costReq model.ShippingCostRequest) (float64, error)
}

type shippingCourierRepository struct {
	DB *gorm.DB
}

func NewShippingCourierRepository(db *gorm.DB) ShippingCourierRepository {
	return &shippingCourierRepository{DB: db}
}

// GetAll implements ShippingCourierRepository.
func (repo *shippingCourierRepository) GetAll() ([]model.ShippingCourier, error) {
	var shippingCouriers []model.ShippingCourier
	err := repo.DB.Find(&shippingCouriers).Error
	if err != nil {
		return nil, err
	}
	return shippingCouriers, nil
}

// GetByID implements ShippingCourierRepository.
func (repo *shippingCourierRepository) GetByID(id int) (*model.ShippingCourier, error) {
	log.Printf("Debug: Get shpping by ID %d", id)
	var shippingCourier model.ShippingCourier
	err := repo.DB.First(&shippingCourier, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &shippingCourier, nil
}

// CalculateShippingCost implements ShippingCourierRepository.
func (repo *shippingCourierRepository) CalculateShippingCost(costReq model.ShippingCostRequest) (float64, error) {
	log.Printf("Debug: Getting distance from third-party API")

	// Construct the URL for the OSRM API
	url := fmt.Sprintf("https://router.project-osrm.org/route/v1/driving/%s;%s?overview=false",
		costReq.OriginLatLong,
		costReq.DestinationLatLong,
	)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error: Failed to fetch data from OSRM API: %v", err)
		return 0, fmt.Errorf("failed to fetch data from OSRM API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: Unexpected response status from OSRM API: %d", resp.StatusCode)
		return 0, fmt.Errorf("unexpected response status from OSRM API: %d", resp.StatusCode)
	}

	// Decode the response body
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Error: Failed to decode response from OSRM API: %v", err)
		return 0, fmt.Errorf("failed to decode response from OSRM API: %w", err)
	}

	// Check if "routes" exists in the response
	routes, ok := result["routes"].([]interface{})
	if !ok || len(routes) == 0 {
		log.Println("Error: No routes found in OSRM API response")
		return 0, errors.New("no routes found in OSRM API response")
	}

	// Extract the distance from the first route
	route, ok := routes[0].(map[string]interface{})
	if !ok {
		log.Println("Error: Invalid route format in OSRM API response")
		return 0, errors.New("invalid route format in OSRM API response")
	}

	distance, ok := route["distance"].(float64)
	if !ok || distance == 0 {
		log.Println("Error: Distance not found or invalid in OSRM API response")
		return 0, errors.New("distance not found or invalid in OSRM API response")
	}

	log.Printf("Debug: Successfully fetched distance: %f meters", distance)
	return distance, nil
}
