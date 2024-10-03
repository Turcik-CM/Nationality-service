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
		Country:     "dodi1",
		City:        "dodi1",
		Nationality: "dodi1",
		ImageUrl:    "dodi1",
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
		Id: "92a7fc4b-bb19-46a9-85d2-7fc9cdac3e05",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestUpdateCountry(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	h := NewCountriesStorage(db)
	res, err := h.UpdateCountry(&pb.UpdateCountryRequest{
		Id:      "1582b2e1-45e8-4786-85cb-9cd219d3b140",
		Country: "dodi222",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestDeleteCountry(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	h := NewCountriesStorage(db)
	res, err := h.DeleteCountry(&pb.DeleteCountryRequest{
		Id: "81bb127c-e323-453a-beb9-1e2bab7ac3b5",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestListCountries(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	h := NewCountriesStorage(db)
	res, err := h.ListCountries(&pb.ListCountriesRequest{
		Limit:   2,
		Country: "dodi222",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
