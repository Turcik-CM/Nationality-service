package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	pb "nationality-service/genproto/nationality"
	"testing"
)

func Connect() (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "dodi", "cmt")

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestCreateCountry(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	h := NewCountriesStorage(db)

	res := pb.CreateCountryRequest{
		Country:     "dodi",
		City:        "dodi",
		Nationality: "dodi",
		ImageUrl:    "dodi",
	}

	req, err := h.CreateCountry(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestGetCountry(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	h := NewCountriesStorage(db)
	res, err := h.GetCountry(&pb.GetCountryRequest{
		Id: "81bb127c-e323-453a-beb9-1e2bab7ac3b5",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
