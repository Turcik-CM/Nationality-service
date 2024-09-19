package service

import (
	"log/slog"
	pb "tourism-service/genproto/tourism"
	"tourism-service/storage"
)

type TourismService struct {
	history storage.HistoryStorage
	logger  *slog.Logger
	pb.UnimplementedTourismServiceServer
}

func NewTourismService(st storage.HistoryStorage) *TourismService {
	return &TourismService{
		history: st,
	}
}
