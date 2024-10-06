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
		INSERT INTO foods (id, food_name, food_type, country_id, description, ingredients, image_url, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	if err := s.db.QueryRow(query, id, in.FoodName, in.FoodType, in.CountryId, in.Description, in.Ingredients, in.ImageUrl, createdAt).Scan(&id); err != nil {
		return nil, err
	}

	return &pb.NationalFoodResponse{
		Id:          id.String(),
		FoodType:    in.FoodType,
		FoodName:    in.FoodName,
		Description: in.Description,
		CountryId:   in.CountryId,
		ImageUrl:    in.ImageUrl,
		CreatedAt:   createdAt.String(),
	}, nil
}

func (s *NationalFoodsStorage) UpdateNationalFood(in *pb.UpdateNationalFood) (*pb.NationalFoodResponse, error) {
	query := `UPDATE foods SET`
	args := []interface{}{}
	argIndex := 1
	updateFields := []string{}

	if in.FoodName != "" {
		updateFields = append(updateFields, fmt.Sprintf("food_name = $%d", argIndex))
		args = append(args, in.FoodName)
		argIndex++
	}
	if in.FoodType != "" {
		updateFields = append(updateFields, fmt.Sprintf("food_type = $%d", argIndex))
		args = append(args, in.FoodType)
		argIndex++
	}
	if in.CountryId != "" {
		updateFields = append(updateFields, fmt.Sprintf("country_id = $%d", argIndex))
		args = append(args, in.CountryId)
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

	query += fmt.Sprintf(" WHERE id = $%d RETURNING id, food_name, food_type, country_id, description, ingredients, image_url, created_at", argIndex)
	args = append(args, in.Id)

	var res pb.NationalFoodResponse
	err := s.db.QueryRowContext(context.Background(), query, args...).Scan(
		&res.Id, &res.FoodName, &res.FoodType, &res.CountryId, &res.Description, &res.Ingredients, &res.ImageUrl, &res.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to update national food: %v", err)
	}

	return &res, nil
}

func (s *NationalFoodsStorage) GetNationalFoodByID(in *pb.NationalFoodId) (*pb.NationalFoodResponse, error) {
	query := `SELECT id, food_name, food_type, country_id, description, ingredients, image_url, created_at FROM foods WHERE id = $1 and deleted_at = 0`

	var food pb.NationalFoodResponse
	if err := s.db.QueryRow(query, in.Id).Scan(&food.Id, &food.FoodName, &food.FoodType, &food.CountryId, &food.Description, &food.Ingredients, &food.ImageUrl, &food.CreatedAt); err != nil {
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
	query := `SELECT COUNT(*) OVER(), id, food_name, food_type, country_id, description, ingredients, image_url, created_at FROM foods WHERE deleted_at = 0`
	var args []interface{}
	argIndex := 1

	if in.CountryId != "" {
		query += fmt.Sprintf(" AND country_id = $%d", argIndex)
		args = append(args, in.CountryId)
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

	var total string
	var foods []*pb.NationalFoodResponse
	for rows.Next() {
		var food pb.NationalFoodResponse
		if err := rows.Scan(&total, &food.Id, &food.FoodName, &food.FoodType, &food.CountryId, &food.Description, &food.Ingredients, &food.ImageUrl, &food.CreatedAt); err != nil {
			return nil, err
		}
		foods = append(foods, &food)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to fetch national foods: %v", err)
	}

	return &pb.NationalFoodListResponse{
		NationalFood: foods,
		Total:        total,
	}, nil
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
