package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	pb "nationality-service/genproto/nationality"
	"strings"
)

type CountriesStorage struct {
	db *sqlx.DB
}

func NewCountriesStorage(db *sqlx.DB) *CountriesStorage {
	return &CountriesStorage{db: db}
}

func (s *CountriesStorage) CreateCountry(in *pb.CreateCountryRequest) (*pb.CreateCountryResponse, error) {
	query := `
        INSERT INTO countries (country, city, nationality, flag)
        VALUES ($1, $2, $3, $4) RETURNING id, country, city, nationality, flag
    `
	var country pb.Country
	err := s.db.QueryRowContext(context.Background(), query, in.Country, in.City, in.Nationality, in.ImageUrl).Scan(
		&country.Id, &country.Country, &country.City, &country.Nationality, &country.ImageUrl,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating country: %v", err)
	}

	return &pb.CreateCountryResponse{
		Id:          country.Id,
		Country:     country.Country,
		City:        country.City,
		Nationality: country.Nationality,
		ImageUrl:    country.ImageUrl,
	}, nil
}

func (s *CountriesStorage) GetCountry(in *pb.GetCountryRequest) (*pb.GetCountryResponse, error) {
	query := `
        SELECT id, country, city, nationality, flag
        FROM countries
        WHERE id = $1
    `
	var country pb.Country
	err := s.db.QueryRowContext(context.Background(), query, in.Id).Scan(
		&country.Id, &country.Country, &country.City, &country.Nationality, &country.ImageUrl,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting country by ID: %v", err)
	}

	return &pb.GetCountryResponse{
		Id:          country.Id,
		Country:     country.Country,
		City:        country.City,
		Nationality: country.Nationality,
		ImageUrl:    country.ImageUrl,
	}, nil
}

func (s *CountriesStorage) UpdateCountry(in *pb.UpdateCountryRequest) (*pb.UpdateCountryResponse, error) {
	query := `UPDATE countries SET`
	args := []interface{}{}
	argIndex := 1
	updateFields := []string{}

	if in.Country != "" {
		updateFields = append(updateFields, fmt.Sprintf("country = $%d", argIndex))
		args = append(args, in.Country)
		argIndex++
	}
	if in.City != "" {
		updateFields = append(updateFields, fmt.Sprintf("city = $%d", argIndex))
		args = append(args, in.City)
		argIndex++
	}
	if in.Nationality != "" {
		updateFields = append(updateFields, fmt.Sprintf("nationality = $%d", argIndex))
		args = append(args, in.Nationality)
		argIndex++
	}
	if in.ImageUrl != "" {
		updateFields = append(updateFields, fmt.Sprintf("flag = $%d", argIndex))
		args = append(args, in.ImageUrl)
		argIndex++
	}

	if len(updateFields) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query += fmt.Sprintf(" %s WHERE id = $%d RETURNING id, country, city, nationality, flag",
		strings.Join(updateFields, ", "), argIndex)
	args = append(args, in.Id)

	var updated pb.Country
	err := s.db.QueryRowContext(context.Background(), query, args...).Scan(
		&updated.Id, &updated.Country, &updated.City, &updated.Nationality, &updated.ImageUrl,
	)
	if err != nil {
		return nil, fmt.Errorf("error updating country: %v", err)
	}

	return &pb.UpdateCountryResponse{
		Id:          updated.Id,
		Country:     updated.Country,
		City:        updated.City,
		Nationality: updated.Nationality,
		ImageUrl:    updated.ImageUrl,
	}, nil
}

func (s *CountriesStorage) DeleteCountry(in *pb.DeleteCountryRequest) (*pb.Message, error) {
	query := `DELETE FROM countries WHERE id = $1`
	_, err := s.db.ExecContext(context.Background(), query, in.Id)
	if err != nil {
		return nil, fmt.Errorf("error deleting country: %v", err)
	}

	return &pb.Message{Message: "Country deleted successfully"}, nil
}

func (s *CountriesStorage) ListCountries(in *pb.ListCountriesRequest) (*pb.ListCountriesResponse, error) {
	query := `
        SELECT id, country, city, nationality, flag
        FROM countries
    `
	if in.Country != "" {
		query += ` WHERE country ILIKE '%' || $1 || '%'`
	}

	if in.City != "" {
		if in.Country != "" {
			query += ` AND city ILIKE '%' || $2 || '%'`
		} else {
			query += ` WHERE city ILIKE '%' || $1 || '%'`
		}
	}

	if in.Limit > 0 {
		query += ` LIMIT $3`
	}

	if in.Offset >= 0 {
		query += ` OFFSET $4`
	}

	rows, err := s.db.QueryContext(context.Background(), query, in.Country, in.City, in.Limit, in.Offset)
	if err != nil {
		return nil, fmt.Errorf("error listing countries: %v", err)
	}
	defer rows.Close()

	countries := make([]*pb.Country, 0)
	for rows.Next() {
		var country pb.Country
		if err := rows.Scan(&country.Id, &country.Country, &country.City, &country.Nationality, &country.ImageUrl); err != nil {
			return nil, fmt.Errorf("error scanning country: %v", err)
		}
		countries = append(countries, &country)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %v", err)
	}

	return &pb.ListCountriesResponse{Countries: countries}, nil
}
