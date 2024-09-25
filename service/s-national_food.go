package service

import (
	"context"
	pb "nationality-service/genproto/nationality"
)

func (s *TourismService) CreateNationalFood(ctx context.Context, in *pb.NationalFood) (*pb.NationalFoodResponse, error) {
	req, err := s.nat_food.CreateNationalFood(in)
	if err != nil {
		s.logger.Error("Erorr in CreateNationalFood", err)
		return nil, err
	}
	return req, nil
}

func (s *TourismService) UpdateNationalFoods(ctx context.Context, in *pb.UpdateNationalFood) (*pb.NationalFoodResponse, error) {
	req, err := s.nat_food.UpdateNationalFood(in)
	if err != nil {
		s.logger.Error("Erorr in UpdateNationalFood", err)
		return nil, err
	}
	return req, nil
}

func (s *TourismService) GetNationalFoodByID(ctx context.Context, in *pb.NationalFoodId) (*pb.NationalFoodResponse, error) {
	req, err := s.nat_food.GetNationalFoodByID(in)
	if err != nil {
		s.logger.Error("Erorr in GetNationalFoodByID", err)
		return nil, err
	}
	return req, nil
}

func (s *TourismService) DeleteNationalFood(ctx context.Context, in *pb.NationalFoodId) (*pb.Message, error) {
	req, err := s.nat_food.DeleteNationalFood(in)
	if err != nil {
		s.logger.Error("Erorr in DeleteNationalFood", err)
		return nil, err
	}
	return req, nil
}

func (s *TourismService) ListNationalFood(ctx context.Context, in *pb.NationalFoodList) (*pb.NationalFoodListResponse, error) {
	req, err := s.nat_food.ListNationalFoods(in)
	if err != nil {
		s.logger.Error("Erorr in ListNationalFood", err)
		return nil, err
	}
	return req, nil
}

func (s *TourismService) AddNationalFoodImage(ctx context.Context, in *pb.NationalFoodImage) (*pb.Message, error) {
	req, err := s.nat_food.AddImageUrll(in)
	if err != nil {
		s.logger.Error("Erorr in AddNationalFoodImage", err)
		return nil, err
	}
	return req, nil
}
