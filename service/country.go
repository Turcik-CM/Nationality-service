package service

import (
	"context"
	"fmt"
	pb "nationality-service/genproto/nationality"
)

// CreateCountry handles the creation of a new country
func (s *TourismService) CreateCountry(ctx context.Context, in *pb.CreateCountryRequest) (*pb.CreateCountryResponse, error) {
	resp, err := s.CreateCountry(ctx, in)
	if err != nil {
		s.logger.Error("Error in CreateCountry: ", err)
		return nil, fmt.Errorf("failed to create country: %v", err)
	}
	return resp, nil
}

// GetCountry retrieves a country by its ID
func (s *TourismService) GetCountry(ctx context.Context, in *pb.GetCountryRequest) (*pb.GetCountryResponse, error) {
	resp, err := s.GetCountry(ctx, in)
	if err != nil {
		s.logger.Error("Error in GetCountry: ", err)
		return nil, fmt.Errorf("failed to get country: %v", err)
	}
	return resp, nil
}

// UpdateCountry updates an existing country
func (s *TourismService) UpdateCountry(ctx context.Context, in *pb.UpdateCountryRequest) (*pb.UpdateCountryResponse, error) {
	resp, err := s.UpdateCountry(ctx, in)
	if err != nil {
		s.logger.Error("Error in UpdateCountry: ", err)
		return nil, fmt.Errorf("failed to update country: %v", err)
	}
	return resp, nil
}

// DeleteCountry deletes a country by its ID
func (s *TourismService) DeleteCountry(ctx context.Context, in *pb.DeleteCountryRequest) (*pb.Message, error) {
	resp, err := s.DeleteCountry(ctx, in)
	if err != nil {
		s.logger.Error("Error in DeleteCountry: ", err)
		return nil, fmt.Errorf("failed to delete country: %v", err)
	}
	return resp, nil
}

// ListCountries retrieves a list of countries based on provided filters
func (s *TourismService) ListCountries(ctx context.Context, in *pb.ListCountriesRequest) (*pb.ListCountriesResponse, error) {
	resp, err := s.ListCountries(ctx, in)
	if err != nil {
		s.logger.Error("Error in ListCountries: ", err)
		return nil, fmt.Errorf("failed to list countries: %v", err)
	}
	return resp, nil
}
