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
		Country:     "Uzbekistan",
		Name:        "dodi",
		Description: "dodi",
		Nationality: "dodi",
		ImageUrl:    "null",
		Rating:      1200000,
		FoodType:    "nimadur",
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
		Id:   "da8b54e8-0675-4b22-8970-46d8af4f3dbf",
		Name: "1111",
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
		Id: "da8b54e8-0675-4b22-8970-46d8af4f3dbf",
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
		Limit:  1,
		Offset: 0,
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
		Id:       "da8b54e8-0675-4b22-8970-46d8af4f3dbf",
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
		Id: "1a5032a0-01b8-4c3b-93b7-726c9e09aab9",
	}
	food := NewNationalFoodsStorage(db)
	req, err := food.DeleteNationalFood(&res)
	if err != nil {
		t.Fatalf("cannot create request: %v", err)
		return
	}
	fmt.Println(req)
}
