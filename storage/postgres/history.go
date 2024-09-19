package postgres

import (
	pb "tourism-service/genproto/tourism"
	"tourism-service/storage"

	"github.com/jmoiron/sqlx"
)

type HistoryRepo struct {
	db *sqlx.DB
}

func NewHistoryRepo(db *sqlx.DB) storage.HistoryStorage {
	return &HistoryRepo{
		db: db,
	}
}

func (h *HistoryRepo) AddHistorical(in *pb.Historical) (*pb.HistoricalResponse, error) {
	return &pb.HistoricalResponse{}, nil
}
func (h *HistoryRepo) UpdateHistorical(in *pb.UpdateAHistorical) (*pb.HistoricalResponse, error) {
	return &pb.HistoricalResponse{}, nil
}
func (h *HistoryRepo) GetHistoricalByID(in *pb.HistoricalId) (*pb.HistoricalResponse, error) {
	return &pb.HistoricalResponse{}, nil
}
func (h *HistoryRepo) ListHistorical(in *pb.HistoricalList) (*pb.HistoricalListResponse, error) {
	return &pb.HistoricalListResponse{}, nil
}
func (h *HistoryRepo) DeleteHistorical(in *pb.HistoricalId) (*pb.Message, error) {
	return &pb.Message{}, nil
}
func (h *HistoryRepo) SearchHistorical(in *pb.HistoricalSearch) (*pb.HistoricalListResponse, error) {
	return &pb.HistoricalListResponse{}, nil
}
func (h *HistoryRepo) GetHistoricalByCountry(in *pb.HistoricalCountry) (*pb.HistoricalListResponse, error) {
	return &pb.HistoricalListResponse{}, nil
}
