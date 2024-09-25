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

type AttractionsStorage struct {
	db *sqlx.DB
}

func NewAttractionsStorage(db *sqlx.DB) *AttractionsStorage {
	return &AttractionsStorage{db: db}
}

// CreateAttraction i nserts a new attraction into the database using Protobuf structures.
func (s *AttractionsStorage) CreateAttraction(in *pb.Attraction) (*pb.AttractionResponse, error) {
	id := uuid.New()
	createdAt := time.Now()

	query := `
		INSERT INTO attractions (id, category, name, description, country, location, image_url, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err := s.db.QueryRowContext(context.Background(), query, id, in.Category, in.Name, in.Description, in.Country, in.Location, in.ImageUrl, createdAt).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("error creating attraction: %v", err)
	}

	return &pb.AttractionResponse{
		Id:          id.String(),
		Category:    in.Category,
		Name:        in.Name,
		Description: in.Description,
		Country:     in.Country,
		Location:    in.Location,
		ImageUrl:    in.ImageUrl,
		CreatedAt:   createdAt.String(),
	}, nil
}

// GetAttractionByID retrieves an attraction by its ID using Protobuf.
func (s *AttractionsStorage) GetAttractionByID(in *pb.AttractionId) (*pb.AttractionResponse, error) {
	query := `
		SELECT id, category, name, description, country, location, image_url, created_at, updated_at
		FROM attractions
		WHERE id = $1 AND deleted_at = 0`

	var attraction pb.AttractionResponse
	err := s.db.QueryRowContext(context.Background(), query, in.Id).Scan(
		&attraction.Id, &attraction.Category, &attraction.Name, &attraction.Description, &attraction.Country,
		&attraction.Location, &attraction.ImageUrl, &attraction.CreatedAt, &attraction.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting attraction by ID: %v", err)
	}

	return &attraction, nil
}

// UpdateAttraction updates an existing attraction using Protobuf.
func (s *AttractionsStorage) UpdateAttraction(in *pb.UpdateAttraction) (*pb.AttractionResponse, error) {
	query := `UPDATE attractions SET`
	args := []interface{}{}
	argIndex := 1
	updateFields := []string{}

	if in.Category != "" {
		updateFields = append(updateFields, fmt.Sprintf("category = $%d", argIndex))
		args = append(args, in.Category)
		argIndex++
	}
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
	if in.Location != "" {
		updateFields = append(updateFields, fmt.Sprintf("location = $%d", argIndex))
		args = append(args, in.Location)
		argIndex++
	}
	if in.ImageUrl != "" {
		updateFields = append(updateFields, fmt.Sprintf("image_url = $%d", argIndex))
		args = append(args, in.ImageUrl)
		argIndex++
	}

	if len(updateFields) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	updateFields = append(updateFields, fmt.Sprintf("updated_at = $%d", argIndex))
	args = append(args, time.Now())
	argIndex++

	query += fmt.Sprintf(" %s WHERE id = $%d RETURNING id, category, name, description, country, location, image_url, created_at, updated_at",
		strings.Join(updateFields, ", "), argIndex)
	args = append(args, in.Id)

	var updated pb.AttractionResponse
	err := s.db.QueryRowContext(context.Background(), query, args...).Scan(
		&updated.Id, &updated.Category, &updated.Name, &updated.Description, &updated.Country,
		&updated.Location, &updated.ImageUrl, &updated.CreatedAt, &updated.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("error updating attraction: %v", err)
	}

	return &updated, nil
}

// DeleteAttraction soft deletes an attraction by setting the deleted_at field.
func (s *AttractionsStorage) DeleteAttraction(in *pb.AttractionId) (*pb.Message, error) {
	query := `update attractions set deleted_at = date_part('epoch', current_timestamp)::INT
                  where id = $1 and deleted_at = 0`

	_, err := s.db.ExecContext(context.Background(), query, in.Id)
	if err != nil {
		return nil, fmt.Errorf("error deleting attraction: %v", err)
	}

	return &pb.Message{Message: "Attraction deleted successfully"}, nil
}

func (s *AttractionsStorage) ListAttractions(in *pb.AttractionList) (*pb.AttractionListResponse, error) {
	query := `
		SELECT id, category, name, description, country, location, image_url, created_at, updated_at
		FROM attractions
		WHERE deleted_at = 0`

	var args []interface{}
	argIndex := 1

	if in.Country != "" {
		query += fmt.Sprintf(" AND country = $%d", argIndex)
		args = append(args, in.Country)
		argIndex++
	}
	if in.Category != "" {
		query += fmt.Sprintf(" AND category = $%d", argIndex)
		args = append(args, in.Category)
		argIndex++
	}
	if in.Name != "" {
		query += fmt.Sprintf(" AND name ILIKE '%%' || $%d || '%%'", argIndex)
		args = append(args, in.Name)
		argIndex++
	}
	if in.Description != "" {
		query += fmt.Sprintf(" AND description ILIKE '%%' || $%d || '%%'", argIndex)
		args = append(args, in.Description)
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
	}

	rows, err := s.db.QueryContext(context.Background(), query, args...)
	if err != nil {
		return nil, fmt.Errorf("error listing attractions: %v", err)
	}
	defer rows.Close()

	var attractions []*pb.AttractionResponse
	for rows.Next() {
		var attraction pb.AttractionResponse
		if err := rows.Scan(&attraction.Id, &attraction.Category, &attraction.Name, &attraction.Description, &attraction.Country, &attraction.Location, &attraction.ImageUrl, &attraction.CreatedAt, &attraction.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning attraction: %v", err)
		}
		attractions = append(attractions, &attraction)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %v", err)
	}

	return &pb.AttractionListResponse{Attractions: attractions}, nil
}

// SearchAttractions searches for attractions by name or description.
func (s *AttractionsStorage) SearchAttractions(in *pb.AttractionSearch) (*pb.AttractionListResponse, error) {
	query := `
		SELECT id, category, name, description, country, location, image_url, created_at, updated_at
		FROM attractions
		WHERE (name ILIKE '%' || $1 || '%' OR description ILIKE '%' || $1 || '%') 
		AND deleted_at = 0
		LIMIT $2 OFFSET $3`

	rows, err := s.db.QueryContext(context.Background(), query, in.SearchTerm, in.Limit, in.Offset)
	if err != nil {
		return nil, fmt.Errorf("error searching attractions: %v", err)
	}
	defer rows.Close()

	var attractions []*pb.AttractionResponse
	for rows.Next() {
		var attraction pb.AttractionResponse
		if err := rows.Scan(&attraction.Id, &attraction.Category, &attraction.Name, &attraction.Description, &attraction.Country, &attraction.Location, &attraction.ImageUrl, &attraction.CreatedAt, &attraction.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning attraction: %v", err)
		}
		attractions = append(attractions, &attraction)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %v", err)
	}

	return &pb.AttractionListResponse{Attractions: attractions}, nil
}

func (s *AttractionsStorage) AddImageUrl(in *pb.AttractionImage) (*pb.Message, error) {
	query := `
		UPDATE attractions SET image_url = $1 WHERE id = $2
	`
	_, err := s.db.Exec(query, in.ImageUrl, in.Id)
	if err != nil {
		return nil, fmt.Errorf("error adding image url: %v", err)
	}

	return &pb.Message{
		Message: "image url added",
	}, nil
}

func (s *AttractionsStorage) RemoveHistoricalImage(in *pb.HistoricalImage) (*pb.Message, error) {
	query := `
		UPDATE attractions SET image_url = $1 WHERE id = $2
	`
	_, err := s.db.Exec(query, "no imag", in.Id)
	if err != nil {
		return nil, fmt.Errorf("error adding image url: %v", err)
	}

	return &pb.Message{
		Message: "no image url removed",
	}, nil
}
