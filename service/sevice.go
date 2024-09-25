package service

import (
	"log/slog"
	pb "nationality-service/genproto/nationality"
	"nationality-service/storage"
)

type TourismService struct {
	history  storage.HistoryStorage
	nat_food storage.NationalFoodsStorage
	att      storage.AttractionsStorage
	logger   *slog.Logger
	pb.UnimplementedNationalityServiceServer
}

func NewTourismService(his storage.HistoryStorage, net_food storage.NationalFoodsStorage, att storage.AttractionsStorage, logger *slog.Logger) *TourismService {
	return &TourismService{
		history:  his,
		nat_food: net_food,
		att:      att,
		logger:   logger,
	}
}
