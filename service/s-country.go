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

func (s *TourismService) CreateCity(ctx context.Context, in *pb.CreateCityRequest) (*pb.CreateCityResponse, error) {
	resp, err := s.country.CreateCity(in)
	if err != nil {
		s.logger.Error("Error in CreateCity: ", err)
		return nil, fmt.Errorf("failed to create city: %v", err)
	}
	return resp, nil
}

func (s *TourismService) GetCity(ctx context.Context, in *pb.GetCityRequest) (*pb.CreateCityResponse, error) {
	resp, err := s.country.GetCity(in)
	if err != nil {
		s.logger.Error("Error in GetCity: ", err)
		return nil, fmt.Errorf("failed to get city: %v", err)
	}
	return resp, nil
}

func (s *TourismService) UpdateCity(ctx context.Context, in *pb.CreateCityResponse) (*pb.CreateCityResponse, error) {
	resp, err := s.country.UpdateCity(in)
	if err != nil {
		s.logger.Error("Error in UpdateCity: ", err)
		return nil, fmt.Errorf("failed to update city: %v", err)
	}
	return resp, nil
}

func (s *TourismService) DeleteCity(ctx context.Context, in *pb.GetCityRequest) (*pb.Message, error) {
	resp, err := s.country.DeleteCity(in)
	if err != nil {
		s.logger.Error("Error in DeleteCity: ", err)
		return nil, fmt.Errorf("failed to delete city: %v", err)
	}
	return resp, nil
}

func (s *TourismService) ListCity(ctx context.Context, in *pb.ListCityRequest) (*pb.ListCityResponse, error) {
	resp, err := s.country.ListCity(in)
	if err != nil {
		s.logger.Error("Error in ListCity: ", err)
		return nil, fmt.Errorf("failed to list city: %v", err)
	}
	return resp, nil
}

func (s *TourismService) GetBYCount(ctx context.Context, in *pb.CountryId) (*pb.GetCountryId, error) {
	resp, err := s.country.GetBYCount(in)
	if err != nil {
		s.logger.Error("Error in GetBYCount: ", err)
		return nil, fmt.Errorf("failed to get by count: %v", err)
	}
	return resp, nil
}
