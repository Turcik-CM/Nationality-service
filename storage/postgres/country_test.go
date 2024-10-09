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
		Name:     "dodi111",
		ImageUrl: "dodi111",
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
		Id: "403967ee-6b97-4e1e-804d-c1d6ef99cd07",
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
		Id: "403967ee-6b97-4e1e-804d-c1d6ef99cd07",
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
		CountryId: "5b74b327-619a-485b-b6d2-0ae1b6ca02a4",
		Name:      "sd111111",
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
		Id: "84179f27-ec53-462d-97b0-a5aaef8c503b",
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
		Id: "537507a1-6557-4a52-97b6-d3fe0cf9293b",
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
