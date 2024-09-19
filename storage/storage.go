package storage

import (
	pb "tourism-service/genproto/tourism"
)

type HistoryStorage interface {
	AddHistorical(in *pb.Historical) (*pb.HistoricalResponse, error)
	UpdateHistorical(in *pb.UpdateAHistorical) (*pb.HistoricalResponse, error)
	GetHistoricalByID(in *pb.HistoricalId) (*pb.HistoricalResponse, error)
	ListHistorical(in *pb.HistoricalList) (*pb.HistoricalListResponse, error)
	DeleteHistorical(in *pb.HistoricalId) (*pb.Message, error)
	SearchHistorical(in *pb.HistoricalSearch) (*pb.HistoricalListResponse, error)
	GetHistoricalByCountry(in *pb.HistoricalCountry) (*pb.HistoricalListResponse, error)
}
