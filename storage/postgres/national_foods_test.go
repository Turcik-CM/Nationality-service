package postgres

import (
	"fmt"
	pb "nationality-service/genproto/nationality"
	"nationality-service/pkg/config"
	"testing"
	"time"
)

func TestCreateNationalFood(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatalf("cannot connect to database: %v", err)
		return
	}
	defer db.Close()
	res := pb.NationalFood{
		FoodName:    "dodi12",
		Description: "dodi12",
		CountryId:   "2140e218-b8fd-4ff0-a5c7-bd18dccffc08",
		ImageUrl:    "null",
		Ingredients: "dodi",
		CreatedAt:   time.Now().String(),
	}
	food := NewNationalFoodsStorage(db)

	req, err := food.CreateNationalFood(&res)
	if err != nil {
		t.Fatalf("cannot create request: %v", err)
		return
	}
	fmt.Println(req)
}

func TestUpdateNationalFood(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatalf("cannot connect to database: %v", err)
		return
	}
	defer db.Close()
	res := pb.UpdateNationalFood{
		Id:       "aebdac5f-9e20-437a-9a69-0b16e7352c42",
		FoodName: "5455",
	}
	food := NewNationalFoodsStorage(db)
	req, err := food.UpdateNationalFood(&res)
	if err != nil {
		t.Fatalf("cannot create request: %v", err)
		return
	}
	fmt.Println(req)
}

func TestGetNationalFoodByID(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatalf("cannot connect to database: %v", err)
		return
	}
	defer db.Close()
	res := pb.NationalFoodId{
		Id: "aebdac5f-9e20-437a-9a69-0b16e7352c42",
	}
	food := NewNationalFoodsStorage(db)

	req, err := food.GetNationalFoodByID(&res)
	if err != nil {
		t.Fatalf("cannot create request: %v", err)
		return
	}
	fmt.Println(req)
}

func TestListNationalFoods(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatalf("cannot connect to database: %v", err)
		return
	}
	defer db.Close()
	res := pb.NationalFoodList{
		//Limit: 1,
		//Offset: 0,
		CountryId: "2140e218-b8fd-4ff0-a5c7-bd18dccffc08",
	}
	foods := NewNationalFoodsStorage(db)
	req, err := foods.ListNationalFoods(&res)
	if err != nil {
		t.Fatalf("cannot create request: %v", err)
		return
	}
	fmt.Println(req)
}

func TestAddImageUrll(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatalf("cannot connect to database: %v", err)
		return
	}
	defer db.Close()

	res := pb.NationalFoodImage{
		Id:       "aebdac5f-9e20-437a-9a69-0b16e7352c42",
		ImageUrl: "rans joylandi",
	}
	food := NewNationalFoodsStorage(db)
	req, err := food.AddImageUrll(&res)
	if err != nil {
		t.Fatalf("cannot create request: %v", err)
		return
	}
	fmt.Println(req)
}

func TestDeleteNationalFood(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatalf("cannot connect to database: %v", err)
		return
	}
	defer db.Close()
	res := pb.NationalFoodId{
		Id: "aebdac5f-9e20-437a-9a69-0b16e7352c42",
	}
	food := NewNationalFoodsStorage(db)
	req, err := food.DeleteNationalFood(&res)
	if err != nil {
		t.Fatalf("cannot create request: %v", err)
		return
	}
	fmt.Println(req)
}
