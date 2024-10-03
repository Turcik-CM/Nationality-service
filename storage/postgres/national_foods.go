package postgres

import (
	"context"
	"fmt"
	pb "nationality-service/genproto/nationality"
	"nationality-service/storage"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type NationalFoodsStorage struct {
	db *sqlx.DB
}

func NewNationalFoodsStorage(db *sqlx.DB) storage.NationalFoodsStorage {
	return &NationalFoodsStorage{
		db: db,
	}
}

func (s *NationalFoodsStorage) CreateNationalFood(in *pb.NationalFood) (*pb.NationalFoodResponse, error) {
	id := uuid.New()
	createdAt := time.Now()

	if in.ImageUrl == "" {
		in.ImageUrl = "no image"
	}

	query := `
		INSERT INTO foods (id, food_name, food_type, nationality, description, ingredients, image_url, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	if err := s.db.QueryRow(query, id, in.Name, in.FoodType, in.Nationality, in.Description, in.Ingredients, in.ImageUrl, createdAt).Scan(&id); err != nil {
		return nil, err
	}

	return &pb.NationalFoodResponse{
		Id:          id.String(),
		Country:     in.Country,
		Name:        in.Name,
		Description: in.Description,
		ImageUrl:    in.ImageUrl,
		Rating:      in.Rating,
		CreatedAt:   createdAt.String(),
	}, nil
}

func (s *NationalFoodsStorage) UpdateNationalFood(in *pb.UpdateNationalFood) (*pb.NationalFoodResponse, error) {
	query := `UPDATE foods SET`
	args := []interface{}{}
	argIndex := 1
	updateFields := []string{}

	if in.Name != "" {
		updateFields = append(updateFields, fmt.Sprintf("food_name = $%d", argIndex))
		args = append(args, in.Name)
		argIndex++
	}
	if in.FoodType != "" {
		updateFields = append(updateFields, fmt.Sprintf("food_type = $%d", argIndex))
		args = append(args, in.FoodType)
		argIndex++
	}
	if in.Nationality != "" {
		updateFields = append(updateFields, fmt.Sprintf("nationality = $%d", argIndex))
		args = append(args, in.Nationality)
		argIndex++
	}
	if in.Description != "" {
		updateFields = append(updateFields, fmt.Sprintf("description = $%d", argIndex))
		args = append(args, in.Description)
		argIndex++
	}
	if in.Ingredients != "" {
		updateFields = append(updateFields, fmt.Sprintf("ingredients = $%d", argIndex))
		args = append(args, in.Ingredients)
		argIndex++
	}
	if in.ImageUrl != "" {
		updateFields = append(updateFields, fmt.Sprintf("image_url = $%d", argIndex))
		args = append(args, in.ImageUrl)
		argIndex++
	}

	if len(updateFields) > 0 {
		query += fmt.Sprintf(" %s, updated_at = $%d", strings.Join(updateFields, ", "), argIndex)
		args = append(args, time.Now())
		argIndex++
	} else {
		return nil, fmt.Errorf("no fields were updated")
	}

	query += fmt.Sprintf(" WHERE id = $%d RETURNING id, food_name, food_type, nationality, description, ingredients, image_url, created_at, updated_at", argIndex)
	args = append(args, in.Id)

	var res pb.NationalFoodResponse
	err := s.db.QueryRowContext(context.Background(), query, args...).Scan(
		&res.Id, &res.Name, &res.FoodType, &res.Nationality, &res.Description, &res.Ingredients, &res.ImageUrl, &res.CreatedAt, &res.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to update national food: %v", err)
	}

	return &res, nil
}

func (s *NationalFoodsStorage) GetNationalFoodByID(in *pb.NationalFoodId) (*pb.NationalFoodResponse, error) {
	query := `SELECT id, food_name, food_type, nationality, description, ingredients, image_url, created_at, updated_at FROM foods WHERE id = $1`

	var food pb.NationalFoodResponse
	if err := s.db.QueryRow(query, in.Id).Scan(&food.Id, &food.Name, &food.FoodType, &food.Nationality, &food.Description, &food.Ingredients, &food.ImageUrl, &food.CreatedAt, &food.UpdatedAt); err != nil {
		return nil, err
	}

	return &food, nil
}

func (s *NationalFoodsStorage) DeleteNationalFood(in *pb.NationalFoodId) (*pb.Message, error) {
	query := `update foods set deleted_at = date_part('epoch', current_timestamp)::INT
                  where id = $1 and deleted_at = 0`

	if _, err := s.db.Exec(query, in.Id); err != nil {
		return nil, err
	}

	return &pb.Message{Message: "Food deleted successfully"}, nil
}

func (s *NationalFoodsStorage) ListNationalFoods(in *pb.NationalFoodList) (*pb.NationalFoodListResponse, error) {
	query := `SELECT id, food_name, food_type, nationality, description, ingredients, image_url, created_at, updated_at FROM foods WHERE deleted_at = 0`
	args := []interface{}{}
	argIndex := 1

	if in.Country != "" {
		query += fmt.Sprintf(" AND nationality = $%d", argIndex)
		args = append(args, in.Country)
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

	var foods []*pb.NationalFoodResponse
	for rows.Next() {
		var food pb.NationalFoodResponse
		if err := rows.Scan(&food.Id, &food.Name, &food.FoodType, &food.Nationality, &food.Description, &food.Ingredients, &food.ImageUrl, &food.CreatedAt, &food.UpdatedAt); err != nil {
			return nil, err
		}
		foods = append(foods, &food)
	}

	return &pb.NationalFoodListResponse{NationalFood: foods}, nil
}
func (s *NationalFoodsStorage) AddImageUrll(in *pb.NationalFoodImage) (*pb.Message, error) {
	query := `UPDATE foods SET image_url = $1 WHERE id = $2`

	res, err := s.db.Exec(query, in.ImageUrl, in.Id)
	if err != nil {
		return nil, err
	}

	if rowsAffected, err := res.RowsAffected(); err != nil {
		return nil, err
	} else if rowsAffected == 0 {
		return nil, fmt.Errorf("food with id=%s not found", in.Id)
	}

	return &pb.Message{Message: "Image added successfully"}, nil
}
