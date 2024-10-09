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
		Name:     "dodi1",
		ImageUrl: "dodi1",
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
		Id: "e3deb218-0e5a-44e1-8c1b-46617202cb8a",
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
		Id:   "8033c910-b040-42cd-a8c2-545171d75303",
		Name: "dodi",
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
		Limit: 2,
		Name:  "dodi",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

//=====================================

func TestCreateCity(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	h := NewCountriesStorage(db)
	res, err := h.CreateCity(&pb.CreateCityRequest{
		CountryId: "8033c910-b040-42cd-a8c2-545171d75303",
		Name:      "sd",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestGetCity(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	h := NewCountriesStorage(db)
	res, err := h.GetCity(&pb.GetCityRequest{
		Id: "197c0af5-5d5b-408e-88b0-dd5b5f2fe728",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestUpdateCity(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	h := NewCountriesStorage(db)
	req, err := h.UpdateCity(&pb.CreateCityResponse{
		Id:   "a313d1aa-bf9e-4d28-a942-15f2a4c674cc",
		Name: "dodi",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestDeleteCity(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	h := NewCountriesStorage(db)
	req, err := h.DeleteCity(&pb.GetCityRequest{
		Id: "a313d1aa-bf9e-4d28-a942-15f2a4c674cc",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestListCity(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	h := NewCountriesStorage(db)
	res, err := h.ListCity(&pb.ListCityRequest{
		Name: "dodi",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestGetBYCount(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	h := NewCountriesStorage(db)
	req, err := h.GetBYCount(&pb.CountryId{
		Id: "8033c910-b040-42cd-a8c2-545171d75303",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}
