package service

import (
	"context"
	pb "nationality-service/genproto/nationality"
)

func (s *TourismService) CreateAttraction(ctx context.Context, in *pb.Attraction) (*pb.AttractionResponse, error) {
	req, err := s.att.CreateAttraction(in)
	if err != nil {
		s.logger.Error("Error in CreateAttraction", err)
		return nil, err
	}
	return req, nil
}

func (s *TourismService) UpdateAttractions(ctx context.Context, in *pb.UpdateAttraction) (*pb.AttractionResponse, error) {
	req, err := s.att.UpdateAttraction(in)
	if err != nil {
		s.logger.Error("Error in UpdateAttractions", err)
		return nil, err
	}
	return req, nil
}

func (s *TourismService) GetAttractionByID(ctx context.Context, in *pb.AttractionId) (*pb.AttractionResponse, error) {
	req, err := s.att.GetAttractionByID(in)
	if err != nil {
		s.logger.Error("Error in GetAttractionByID", err)
		return nil, err
	}
	return req, nil
}

func (s *TourismService) ListAttraction(ctx context.Context, in *pb.AttractionList) (*pb.AttractionListResponse, error) {
	req, err := s.att.ListAttractions(in)
	if err != nil {
		s.logger.Error("Error in ListAttractions", err)
		return nil, err
	}
	return req, nil
}

func (s *TourismService) DeleteAttraction(ctx context.Context, in *pb.AttractionId) (*pb.Message, error) {
	req, err := s.att.DeleteAttraction(in)
	if err != nil {
		s.logger.Error("Error in DeleteAttraction", err)
		return nil, err
	}
	return req, nil
}

func (s *TourismService) AddAttractionImage(ctx context.Context, in *pb.AttractionImage) (*pb.Message, error) {
	req, err := s.att.AddImageUrl(in)
	if err != nil {
		s.logger.Error("Error in AddAttractionImage", err)
		return nil, err
	}
	return req, nil
}

func (s *TourismService) SearchAttraction(ctx context.Context, in *pb.AttractionSearch) (*pb.AttractionListResponse, error) {
	req, err := s.att.SearchAttractions(in)
	if err != nil {
		s.logger.Error("Error in SearchAttractions", err)
		return nil, err
	}
	return req, nil
}

func (s *TourismService) RemoveHistoricalImage(ctx context.Context, in *pb.HistoricalImage) (*pb.Message, error) {
	req, err := s.att.RemoveHistoricalImage(in)
	if err != nil {
		s.logger.Error("Error in RemoveHistoricalImage", err)
		return nil, err
	}
	return req, nil
}
func (s *TourismService) CreateAttractionType(ctx context.Context, in *pb.Attraction) (*pb.AttractionResponse, error) {
	req, err := s.att.CreateAttraction(in)
	if err != nil {
		s.logger.Error("Error in CreateAttraction", err)
		return nil, err
	}
	return req, nil
}

func (s *TourismService) UpdateAttractionType(ctx context.Context, in *pb.UpdateAttraction) (*pb.AttractionResponse, error) {
	req, err := s.att.UpdateAttraction(in)
	if err != nil {
		s.logger.Error("Error in UpdateAttraction", err)
		return nil, err
	}
	return req, nil
}

func (s *TourismService) GetAttractionTypeByID(ctx context.Context, in *pb.AttractionId) (*pb.AttractionResponse, error) {
	req, err := s.att.GetAttractionByID(in)
	if err != nil {
		s.logger.Error("Error in GetAttractionByID", err)
		return nil, err
	}
	return req, nil
}

func (s *TourismService) ListAttractionTypes(ctx context.Context, in *pb.AttractionList) (*pb.AttractionListResponse, error) {
	req, err := s.att.ListAttractions(in)
	if err != nil {
		s.logger.Error("Error in ListAttractions", err)
		return nil, err
	}
	return req, nil
}

func (s *TourismService) DeleteAttractionType(ctx context.Context, in *pb.AttractionId) (*pb.Message, error) {
	req, err := s.att.DeleteAttraction(in)
	if err != nil {
		s.logger.Error("Error in DeleteAttraction", err)
		return nil, err
	}
	return req, nil
}
