package postgres

import (
	"fmt"
	pb "nationality-service/genproto/nationality"
	"nationality-service/pkg/config"
	"testing"
	"time"
)

func TestAddHistorical(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()

	res := pb.Historical{
		City:        "dodi",
		Name:        "dodi",
		Description: "dodi",
		ImageUrl:    "null",
		CreatedAt:   time.Now().String(),
	}

	his := NewHistoryStorage(db)

	req, err := his.AddHistorical(&res)
	if err != nil {
		t.Errorf("Failed to add history: %v", err)
		return
	}
	fmt.Println(req)
}

func TestUpdateHistoricals(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()
	res := pb.UpdateHistorical{
		Id:          "fc8448fd-11d6-4f0a-bc58-077fb3b94be9",
		Description: "nimadur",
	}

	his := NewHistoryStorage(db)
	req, err := his.UpdateHistoricals(&res)
	if err != nil {
		t.Errorf("Failed to add history: %v", err)
		return
	}
	fmt.Println(req)
}

func TestGetHistoricalByID(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()
	res := pb.HistoricalId{
		Id: "cc39cd85-20be-4a7b-9aaa-065f9e7d5ae5",
	}
	his := NewHistoryStorage(db)
	req, err := his.GetHistoricalByID(&res)
	if err != nil {
		t.Errorf("Failed to add history: %v", err)
		return
	}
	fmt.Println(req)
}

func TestListHistorical(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()
	res := pb.HistoricalList{
		Limit: 1,
		City:  "dodi",
	}
	his := NewHistoryStorage(db)
	req, err := his.ListHistorical(&res)
	if err != nil {
		t.Errorf("Failed to add history: %v", err)
		return
	}
	fmt.Println(req)
}

func TestSearchHistorical(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()
	res := pb.HistoricalSearch{
		Search: "dodi",
	}
	his := NewHistoryStorage(db)
	req, err := his.SearchHistorical(&res)
	if err != nil {
		t.Errorf("Failed to add history: %v", err)
		return
	}
	fmt.Println(req)
}

func TestAddHistoricalImage(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()
	res := pb.HistoricalImage{
		Id:  "cc39cd85-20be-4a7b-9aaa-065f9e7d5ae5",
		Url: "dodi",
	}
	his := NewHistoryStorage(db)
	req, err := his.AddHistoricalImage(&res)
	if err != nil {
		t.Errorf("Failed to add history: %v", err)
		return
	}
	fmt.Println(req)
}

func TestDeleteHistorical(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()
	res := pb.HistoricalId{
		Id: "cc39cd85-20be-4a7b-9aaa-065f9e7d5ae5",
	}
	his := NewHistoryStorage(db)
	req, err := his.DeleteHistorical(&res)
	if err != nil {
		t.Errorf("Failed to add history: %v", err)
		return
	}
	fmt.Println(req)
}
