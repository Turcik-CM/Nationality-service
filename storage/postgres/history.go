package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	pb "nationality-service/genproto/nationality"
)

type HistoryStorage struct {
	db *sqlx.DB
}

func NewHistoryStorage(db *sqlx.DB) *HistoryStorage {
	return &HistoryStorage{db: db}
}

func (s *HistoryStorage) AddHistorical(in *pb.Historical) (*pb.HistoricalResponse, error) {
	id := uuid.New()
	createdAt := time.Now()
	updatedAt := createdAt

	query := `
		INSERT INTO history (id, name, description, country, image_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	if err := s.db.QueryRow(query, id, in.Name, in.Description, in.Country, in.ImageUrl, createdAt, updatedAt).Scan(&id); err != nil {
		return nil, err
	}

	return &pb.HistoricalResponse{
		Id:          id.String(),
		Country:     in.Country,
		City:        in.City,
		Name:        in.Name,
		Description: in.Description,
		ImageUrl:    in.ImageUrl,
		CreatedAt:   createdAt.String(),
		UpdatedAt:   updatedAt.String(),
	}, nil
}

func (s *HistoryStorage) UpdateHistoricals(in *pb.UpdateHistorical) (*pb.HistoricalResponse, error) {
	query := `UPDATE history SET`
	args := []interface{}{}
	argIndex := 1
	updateFields := []string{}

	if in.Name != "" {
		updateFields = append(updateFields, fmt.Sprintf("name = $%d", argIndex))
		args = append(args, in.Name)
		argIndex++
	}
	if in.Description != "" {
		updateFields = append(updateFields, fmt.Sprintf("description = $%d", argIndex))
		args = append(args, in.Description)
		argIndex++
	}
	if in.Country != "" {
		updateFields = append(updateFields, fmt.Sprintf("country = $%d", argIndex))
		args = append(args, in.Country)
		argIndex++
	}
	if in.ImageUrl != "" {
		updateFields = append(updateFields, fmt.Sprintf("image_url = $%d", argIndex))
		args = append(args, in.ImageUrl)
		argIndex++
	}

	if len(updateFields) > 0 {
		query += fmt.Sprintf("%s, updated_at = $%d", strings.Join(updateFields, ", "), argIndex)
		args = append(args, time.Now())
		argIndex++
	} else {
		return nil, fmt.Errorf("no fields were updated")
	}

	query += fmt.Sprintf(" WHERE id = $%d RETURNING id, name, description, country, image_url, created_at, updated_at", argIndex)
	args = append(args, in.Id)

	var res pb.HistoricalResponse
	err := s.db.QueryRowContext(context.Background(), query, args...).Scan(
		&res.Id, &res.Name, &res.Description, &res.Country, &res.ImageUrl, &res.CreatedAt, &res.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *HistoryStorage) GetHistoricalByID(in *pb.HistoricalId) (*pb.HistoricalResponse, error) {
	query := `SELECT id, name, description, country, image_url, created_at, updated_at FROM history WHERE id = $1`

	var historical pb.HistoricalResponse
	if err := s.db.QueryRow(query, in.Id).Scan(&historical.Id, &historical.Name, &historical.Description, &historical.Country, &historical.ImageUrl, &historical.CreatedAt, &historical.UpdatedAt); err != nil {
		return nil, err
	}

	return &historical, nil
}

func (s *HistoryStorage) DeleteHistorical(in *pb.HistoricalId) (*pb.Message, error) {
	query := `DELETE FROM history WHERE id = $1`

	if _, err := s.db.Exec(query, in.Id); err != nil {
		return nil, err
	}

	return &pb.Message{Message: "Historical entry deleted successfully"}, nil
}

func (s *HistoryStorage) ListHistorical(in *pb.HistoricalList) (*pb.HistoricalListResponse, error) {
	query := `SELECT id, name, description, country, image_url, created_at, updated_at FROM history WHERE 1=1`
	args := []interface{}{}
	argIndex := 1

	if in.Country != "" {
		query += fmt.Sprintf(" AND country = $%d", argIndex)
		args = append(args, in.Country)
		argIndex++
	}

	if in.City != "" {
		query += fmt.Sprintf(" AND city = $%d", argIndex)
		args = append(args, in.City)
		argIndex++
	}

	if in.Limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", argIndex)
		args = append(args, in.Limit)
		argIndex++
	}

	if in.Offset >= 0 {
		query += fmt.Sprintf(" OFFSET $%d", argIndex)
		args = append(args, in.Offset)
		argIndex++
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var historicals []*pb.HistoricalResponse
	for rows.Next() {
		var historical pb.HistoricalResponse
		if err := rows.Scan(&historical.Id, &historical.Name, &historical.Description, &historical.Country, &historical.ImageUrl, &historical.CreatedAt, &historical.UpdatedAt); err != nil {
			return nil, err
		}
		historicals = append(historicals, &historical)
	}

	return &pb.HistoricalListResponse{Historical: historicals}, nil
}

func (s *HistoryStorage) SearchHistorical(in *pb.HistoricalSearch) (*pb.HistoricalListResponse, error) {
	query := `SELECT id, name, description, country, image_url, created_at, updated_at FROM history WHERE name ILIKE '%' || $1 || '%' or description ILIKE '%' || $1 || '%'`
	args := []interface{}{in.Search}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var historicals []*pb.HistoricalResponse
	for rows.Next() {
		var historical pb.HistoricalResponse
		if err := rows.Scan(&historical.Id, &historical.Name, &historical.Description, &historical.Country, &historical.ImageUrl, &historical.CreatedAt, &historical.UpdatedAt); err != nil {
			return nil, err
		}
		historicals = append(historicals, &historical)
	}

	return &pb.HistoricalListResponse{Historical: historicals}, nil
}
func (s *HistoryStorage) AddHistoricalImage(in *pb.HistoricalImage) (*pb.Message, error) {
	query := `UPDATE history SET image_url = $1 WHERE id = $2`

	if _, err := s.db.Exec(query, in.Url, in.Id); err != nil {
		return nil, err
	}

	return &pb.Message{Message: "Image added successfully"}, nil
}
