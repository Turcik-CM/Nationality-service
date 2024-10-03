package service

import (
	"context"
	"fmt"
	pb "nationality-service/genproto/nationality"
)

func (s *TourismService) CreateCountry(ctx context.Context, in *pb.CreateCountryRequest) (*pb.CreateCountryResponse, error) {
	resp, err := s.country.CreateCountry(in)
	if err != nil {
		s.logger.Error("Error in CreateCountry: ", err)
		return nil, fmt.Errorf("failed to create country: %v", err)
	}
	return resp, nil
}

func (s *TourismService) GetCountry(ctx context.Context, in *pb.GetCountryRequest) (*pb.GetCountryResponse, error) {
	resp, err := s.country.GetCountry(in)
	if err != nil {
		s.logger.Error("Error in GetCountry: ", err)
		return nil, fmt.Errorf("failed to get country: %v", err)
	}
	return resp, nil
}

func (s *TourismService) UpdateCountry(ctx context.Context, in *pb.UpdateCountryRequest) (*pb.UpdateCountryResponse, error) {
	resp, err := s.country.UpdateCountry(in)
	if err != nil {
		s.logger.Error("Error in UpdateCountry: ", err)
		return nil, fmt.Errorf("failed to update country: %v", err)
	}
	return resp, nil
}

func (s *TourismService) DeleteCountry(ctx context.Context, in *pb.DeleteCountryRequest) (*pb.Message, error) {
	resp, err := s.country.DeleteCountry(in)
	if err != nil {
		s.logger.Error("Error in DeleteCountry: ", err)
		return nil, fmt.Errorf("failed to delete country: %v", err)
	}
	return resp, nil
}

func (s *TourismService) ListCountries(ctx context.Context, in *pb.ListCountriesRequest) (*pb.ListCountriesResponse, error) {
	resp, err := s.country.ListCountries(in)
	if err != nil {
		s.logger.Error("Error in ListCountries: ", err)
		return nil, fmt.Errorf("failed to list countries: %v", err)
	}
	return resp, nil
}
