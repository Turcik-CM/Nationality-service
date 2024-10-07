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
	fmt.Println("dodi")

	query := `
        INSERT INTO countries (name, flag)
        VALUES ($1, $2) RETURNING id, name, flag
    `
	var country pb.Country
	err := s.db.QueryRowContext(context.Background(), query, in.Name, in.ImageUrl).Scan(
		&country.Id, &country.Name, &country.ImageUrl,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating country: %v", err)
	}

	return &pb.CreateCountryResponse{
		Id:       country.Id,
		Name:     country.Name,
		ImageUrl: country.ImageUrl,
	}, nil
}

func (s *CountriesStorage) GetCountry(in *pb.GetCountryRequest) (*pb.GetCountryResponse, error) {
	query := `
        SELECT id, name, flag
        FROM countries
        WHERE id = $1
    `
	var country pb.Country
	err := s.db.QueryRowContext(context.Background(), query, in.Id).Scan(
		&country.Id, &country.Name, &country.ImageUrl,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting country by ID: %v", err)
	}

	return &pb.GetCountryResponse{
		Id:       country.Id,
		Name:     country.Name,
		ImageUrl: country.ImageUrl,
	}, nil
}

func (s *CountriesStorage) UpdateCountry(in *pb.UpdateCountryRequest) (*pb.UpdateCountryResponse, error) {
	query := `UPDATE countries SET`
	args := []interface{}{}
	argIndex := 1
	updateFields := []string{}

	if in.Name != "" {
		updateFields = append(updateFields, fmt.Sprintf("name = $%d", argIndex))
		args = append(args, in.Name)
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

	query += fmt.Sprintf(" %s WHERE id = $%d RETURNING id, name, flag",
		strings.Join(updateFields, ", "), argIndex)
	args = append(args, in.Id)

	var updated pb.Country
	err := s.db.QueryRowContext(context.Background(), query, args...).Scan(
		&updated.Id, &updated.Name, &updated.ImageUrl,
	)
	if err != nil {
		return nil, fmt.Errorf("error updating country: %v", err)
	}

	return &pb.UpdateCountryResponse{
		Id:       updated.Id,
		Name:     updated.Name,
		ImageUrl: updated.ImageUrl,
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
        SELECT COUNT(*) OVER(), id, name, flag
        FROM countries where 1=1
    `
	args := []interface{}{}
	argIndex := 1

	if in.Name != "" {
		query += fmt.Sprintf(" AND name ILIKE '%%' || $%d || '%%'", argIndex)
		args = append(args, in.Name)
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
		return nil, fmt.Errorf("error listing countries: %v", err)
	}
	defer rows.Close()

	var total string
	var countries []*pb.Country
	for rows.Next() {
		var country pb.Country
		if err := rows.Scan(&total, &country.Id, &country.Name, &country.ImageUrl); err != nil {
			return nil, fmt.Errorf("error scanning country: %v", err)
		}
		countries = append(countries, &country)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %v", err)
	}

	return &pb.ListCountriesResponse{
		Countries: countries,
		Total:     total,
	}, nil
}

//===========================================================

func (s *CountriesStorage) CreateCity(in *pb.CreateCityRequest) (*pb.CreateCityResponse, error) {
	fmt.Println("dodi")

	query := `
        INSERT INTO cities (country_id, name)
        VALUES ($1, $2) RETURNING id, country_id, name
    `
	var country pb.CreateCityResponse
	err := s.db.QueryRowContext(context.Background(), query, in.CountryId, in.Name).Scan(
		&country.Id, &country.CountryId, &country.Name,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating country: %v", err)
	}

	return &pb.CreateCityResponse{
		Id:        country.Id,
		Name:      country.Name,
		CountryId: country.CountryId,
	}, nil
}
func (s *CountriesStorage) GetCity(in *pb.GetCityRequest) (*pb.CreateCityResponse, error) {
	query := `
        SELECT id, country_id, name
        FROM cities
        WHERE id = $1
    `
	var country pb.CreateCityResponse
	err := s.db.QueryRowContext(context.Background(), query, in.Id).Scan(
		&country.Id, &country.CountryId, &country.Name,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting country by ID: %v", err)
	}

	return &pb.CreateCityResponse{
		Id:        country.Id,
		Name:      country.Name,
		CountryId: country.CountryId,
	}, nil
}
func (s *CountriesStorage) UpdateCity(in *pb.CreateCityResponse) (*pb.CreateCityResponse, error) {
	query := `UPDATE cities SET`
	args := []interface{}{}
	argIndex := 1
	updateFields := []string{}

	if in.Name != "" {
		updateFields = append(updateFields, fmt.Sprintf("name = $%d", argIndex))
		args = append(args, in.Name)
		argIndex++
	}

	if len(updateFields) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query += fmt.Sprintf(" %s WHERE id = $%d RETURNING id, country_id, name",
		strings.Join(updateFields, ", "), argIndex)
	args = append(args, in.Id)

	var updated pb.CreateCityResponse
	err := s.db.QueryRowContext(context.Background(), query, args...).Scan(
		&updated.Id, &updated.CountryId, &updated.Name,
	)
	if err != nil {
		return nil, fmt.Errorf("error updating country: %v", err)
	}

	return &pb.CreateCityResponse{
		Id:        updated.Id,
		Name:      updated.Name,
		CountryId: updated.CountryId,
	}, nil
}
func (s *CountriesStorage) DeleteCity(in *pb.GetCityRequest) (*pb.Message, error) {
	query := `DELETE FROM cities WHERE id = $1`
	_, err := s.db.ExecContext(context.Background(), query, in.Id)
	if err != nil {
		return nil, fmt.Errorf("error deleting country: %v", err)
	}

	return &pb.Message{Message: "Country deleted successfully"}, nil
}
func (s *CountriesStorage) ListCity(in *pb.ListCityRequest) (*pb.ListCityResponse, error) {
	query := `
        SELECT  COUNT(*) OVER(), id, country_id, name
        FROM cities where 1=1
    `
	args := []interface{}{}
	argIndex := 1

	if in.Name != "" {
		query += fmt.Sprintf(" AND name ILIKE '%%' || $%d || '%%'", argIndex)
		args = append(args, in.Name)
		argIndex++
	}
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
		return nil, fmt.Errorf("error listing countries: %v", err)
	}
	defer rows.Close()

	var total string
	var countries []*pb.CreateCityResponse
	for rows.Next() {
		var country pb.CreateCityResponse
		if err := rows.Scan(&total, &country.Id, &country.CountryId, &country.Name); err != nil {
			return nil, fmt.Errorf("error scanning country: %v", err)
		}
		countries = append(countries, &country)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %v", err)
	}

	return &pb.ListCityResponse{
		Countries: countries,
		Total:     total,
	}, nil
}

func (s *CountriesStorage) GetBYCount(in *pb.CountryId) (*pb.GetCountryId, error) {
	query := `
        SELECT c.id, c.name, cn.name, cn.flag
        FROM cities AS c LEFT OUTER JOIN countries AS cn ON c.country_id = c.id
        WHERE c.country_id = $1
    `
	var countries []*pb.CreateResponse

	rows, err := s.db.QueryContext(context.Background(), query, in.Id)
	if err != nil {
		return nil, fmt.Errorf("error querying countries by country ID: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var country pb.CreateResponse

		err := rows.Scan(&country.Id, &country.Id, &country.CityName, &country.CountryName, &country.FlagUrl)
		if err != nil {
			return nil, fmt.Errorf("error scanning country row: %v", err)
		}
		countries = append(countries, &country)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over country rows: %v", err)
	}

	return &pb.GetCountryId{
		Countries: countries,
	}, nil
}
