package service

import (
	pb "tourism-service/genproto/tourism"
)

func (h *TourismService) AddHistorical(in *pb.Historical) (*pb.HistoricalResponse, error) {
	return &pb.HistoricalResponse{}, nil
}
func (h *TourismService) UpdateHistorical(in *pb.UpdateAHistorical) (*pb.HistoricalResponse, error) {
	return &pb.HistoricalResponse{}, nil
}
func (h *TourismService) GetHistoricalByID(in *pb.HistoricalId) (*pb.HistoricalResponse, error) {
	return &pb.HistoricalResponse{}, nil
}
func (h *TourismService) ListHistorical(in *pb.HistoricalList) (*pb.HistoricalListResponse, error) {
	return &pb.HistoricalListResponse{}, nil
}
func (h *TourismService) DeleteHistorical(in *pb.HistoricalId) (*pb.Message, error) {
	return &pb.Message{}, nil
}
func (h *TourismService) SearchHistorical(in *pb.HistoricalSearch) (*pb.HistoricalListResponse, error) {
	return &pb.HistoricalListResponse{}, nil
}
func (h *TourismService) GetHistoricalByCountry(in *pb.HistoricalCountry) (*pb.HistoricalListResponse, error) {
	return &pb.HistoricalListResponse{}, nil
}
