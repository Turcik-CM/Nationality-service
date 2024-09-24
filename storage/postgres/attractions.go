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

// CreateAttraction inserts a new attraction into the database using Protobuf structures.
func (s *AttractionsStorage) CreateAttraction(ctx context.Context, in *pb.Attraction) (*pb.AttractionResponse, error) {
	id := uuid.New()
	createdAt := time.Now()
	updatedAt := createdAt

	query := `
		INSERT INTO attractions (id, category, name, description, country, location, image_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	err := s.db.QueryRowContext(ctx, query, id, in.Category, in.Name, in.Description, in.Country, in.Location, in.ImageUrl, createdAt, updatedAt).Scan(&id)
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
		UpdatedAt:   updatedAt.String(),
	}, nil
}

// GetAttractionByID retrieves an attraction by its ID using Protobuf.
func (s *AttractionsStorage) GetAttractionByID(ctx context.Context, in *pb.AttractionId) (*pb.AttractionResponse, error) {
	query := `
		SELECT id, category, name, description, country, location, image_url, created_at, updated_at
		FROM attractions
		WHERE id = $1 AND deleted_at IS NULL`

	var attraction pb.AttractionResponse
	err := s.db.QueryRowContext(ctx, query, in.Id).Scan(
		&attraction.Id, &attraction.Category, &attraction.Name, &attraction.Description, &attraction.Country,
		&attraction.Location, &attraction.ImageUrl, &attraction.CreatedAt, &attraction.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting attraction by ID: %v", err)
	}

	return &attraction, nil
}

// UpdateAttraction updates an existing attraction using Protobuf.
func (s *AttractionsStorage) UpdateAttraction(ctx context.Context, in *pb.UpdateAttraction) (*pb.AttractionResponse, error) {
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

	query += fmt.Sprintf("%s, updated_at = $%d WHERE id = $%d RETURNING id, category, name, description, country, location, image_url, created_at, updated_at", strings.Join(updateFields, ", "), argIndex, argIndex+1)
	args = append(args, time.Now(), in.Id)

	var updated pb.AttractionResponse
	err := s.db.QueryRowContext(ctx, query, args...).Scan(
		&updated.Id, &updated.Category, &updated.Name, &updated.Description, &updated.Country,
		&updated.Location, &updated.ImageUrl, &updated.CreatedAt, &updated.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("error updating attraction: %v", err)
	}

	return &updated, nil
}

// DeleteAttraction soft deletes an attraction by setting the deleted_at field.
func (s *AttractionsStorage) DeleteAttraction(ctx context.Context, in *pb.AttractionId) (*pb.Message, error) {
	query := `UPDATE attractions SET deleted_at = $1 WHERE id = $2`

	_, err := s.db.ExecContext(ctx, query, time.Now(), in.Id)
	if err != nil {
		return nil, fmt.Errorf("error deleting attraction: %v", err)
	}

	return &pb.Message{Message: "Attraction deleted successfully"}, nil
}

// ListAttractions retrieves attractions by country with optional filters and pagination using Protobuf.
func (s *AttractionsStorage) ListAttractions(ctx context.Context, in *pb.AttractionList) (*pb.AttractionListResponse, error) {
	// Base query for selecting attractions
	query := `
		SELECT id, category, name, description, country, location, image_url, created_at, updated_at
		FROM attractions
		WHERE deleted_at IS NULL`

	// Arguments slice and filters
	var args []interface{}
	argIndex := 1

	// Dynamically build filters based on input
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

	// Add pagination (LIMIT and OFFSET)
	if in.Limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", argIndex)
		args = append(args, in.Limit)
		argIndex++
	}
	if in.Offset >= 0 {
		query += fmt.Sprintf(" OFFSET $%d", argIndex)
		args = append(args, in.Offset)
	}

	// Execute the query with filters
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error listing attractions: %v", err)
	}
	defer rows.Close()

	// Parse results into protobuf response
	var attractions []*pb.AttractionResponse
	for rows.Next() {
		var attraction pb.AttractionResponse
		if err := rows.Scan(&attraction.Id, &attraction.Category, &attraction.Name, &attraction.Description, &attraction.Country, &attraction.Location, &attraction.ImageUrl, &attraction.CreatedAt, &attraction.UpdatedAt); err != nil {
			return nil, err
		}
		attractions = append(attractions, &attraction)
	}

	// Return the response
	return &pb.AttractionListResponse{Attractions: attractions}, nil
}

// SearchAttractions searches for attractions by name or description.
func (s *AttractionsStorage) SearchAttractions(ctx context.Context, in *pb.AttractionSearch) (*pb.AttractionListResponse, error) {
	query := `
		SELECT id, category, name, description, country, location, image_url, created_at, updated_at
		FROM attractions
		WHERE (name ILIKE '%' || $1 || '%' OR description ILIKE '%' || $1 || '%') AND deleted_at IS NULL
		LIMIT $2 OFFSET $3`

	rows, err := s.db.QueryContext(ctx, query, in.SearchTerm, in.Limit, in.Offset)
	if err != nil {
		return nil, fmt.Errorf("error searching attractions: %v", err)
	}
	defer rows.Close()

	var attractions []*pb.AttractionResponse
	for rows.Next() {
		var attraction pb.AttractionResponse
		if err := rows.Scan(&attraction.Id, &attraction.Category, &attraction.Name, &attraction.Description, &attraction.Country, &attraction.Location, &attraction.ImageUrl, &attraction.CreatedAt, &attraction.UpdatedAt); err != nil {
			return nil, err
		}
		attractions = append(attractions, &attraction)
	}

	return &pb.AttractionListResponse{Attractions: attractions}, nil
}

// AddImageUrl adds a url to the table
func (s *AttractionsStorage) AddImageUrl(ctx context.Context, in *pb.AttractionImage) (*pb.AttractionResponse, error) {
	query := `
		UPDATE attractions SET image_url = $1 WHERE id = $2 RETURNING id, category, name, description, country, location, image_url, created_at, updated_at
	`

	var attraction pb.AttractionResponse
	err := s.db.QueryRowContext(ctx, query, in.ImageUrl, in.Id).Scan(
		&attraction.Id, &attraction.Category, &attraction.Name, &attraction.Description, &attraction.Country,
		&attraction.Location, &attraction.ImageUrl, &attraction.CreatedAt, &attraction.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("error adding image url: %v", err)
	}

	return &attraction, nil
}
