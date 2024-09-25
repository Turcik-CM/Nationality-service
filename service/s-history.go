package service

import (
	"context"
	pb "nationality-service/genproto/nationality"
)

func (s *TourismService) AddHistorical(ctx context.Context, in *pb.Historical) (*pb.HistoricalResponse, error) {
	res, err := s.history.AddHistorical(in)
	if err != nil {
		s.logger.Error("Error in", "err", err)
		return nil, err
	}
	return res, nil
}

func (s *TourismService) UpdateHistoricals(ctx context.Context, in *pb.UpdateHistorical) (*pb.HistoricalResponse, error) {
	res, err := s.history.UpdateHistoricals(in)
	if err != nil {
		s.logger.Error("Error in", "err", err)
		return nil, err
	}
	return res, nil
}

func (s *TourismService) GetHistoricalByID(ctx context.Context, in *pb.HistoricalId) (*pb.HistoricalResponse, error) {
	res, err := s.history.GetHistoricalByID(in)
	if err != nil {
		s.logger.Error("Error in", "err", err)
		return nil, err
	}
	return res, nil
}

func (s *TourismService) DeleteHistorical(ctx context.Context, in *pb.HistoricalId) (*pb.Message, error) {
	res, err := s.history.DeleteHistorical(in)
	if err != nil {
		s.logger.Error("Error in", "err", err)
		return nil, err
	}
	return res, nil
}

func (s *TourismService) ListHistorical(ctx context.Context, in *pb.HistoricalList) (*pb.HistoricalListResponse, error) {
	res, err := s.history.ListHistorical(in)
	if err != nil {
		s.logger.Error("Error in", "err", err)
		return nil, err
	}
	return res, nil
}

func (s *TourismService) SearchHistorical(ctx context.Context, in *pb.HistoricalSearch) (*pb.HistoricalListResponse, error) {
	res, err := s.history.SearchHistorical(in)
	if err != nil {
		s.logger.Error("Error in", "err", err)
		return nil, err
	}
	return res, nil
}

func (s *TourismService) AddHistoricalImage(ctx context.Context, in *pb.HistoricalImage) (*pb.Message, error) {
	res, err := s.history.AddHistoricalImage(in)
	if err != nil {
		s.logger.Error("Error in", "err", err)
		return nil, err
	}
	return res, nil
}
