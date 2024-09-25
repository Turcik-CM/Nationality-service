package postgres

import (
	"fmt"
	pb "nationality-service/genproto/nationality"
	"nationality-service/pkg/config"
	"testing"
	"time"
)

func TestCreateAttraction(t *testing.T) {
	cfg := config.Load()

	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()
	res := pb.Attraction{
		Country:     "Uzbekistan",
		Name:        "dodi",
		Description: "dodi",
		Category:    "culture",
		Location:    "dodi",
		ImageUrl:    "dodi",
		CreatedAt:   time.Now().String(),
	}

	atts := NewAttractionsStorage(db)

	req, err := atts.CreateAttraction(&res)
	if err != nil {
		t.Errorf("Failed to create attraction: %v", err)
		return
	}
	fmt.Println(req)
}

func TestGetAttractionByID(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()
	atts := NewAttractionsStorage(db)

	res := pb.AttractionId{
		Id: "fd2ceee5-e9c9-42d1-9f35-2488d3fb3f30",
	}

	req, err := atts.GetAttractionByID(&res)
	if err != nil {
		t.Errorf("Failed to get attraction: %v", err)
		return
	}
	fmt.Println(req)
}

func TestUpdateAttraction(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()
	atts := NewAttractionsStorage(db)
	res := pb.UpdateAttraction{
		Id:   "fd2ceee5-e9c9-42d1-9f35-2488d3fb3f30",
		Name: "dodi dodi",
	}
	req, err := atts.UpdateAttraction(&res)
	if err != nil {
		t.Errorf("Failed to update attraction: %v", err)
		return
	}
	fmt.Println(req)
}

func TestListAttractions(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()
	atts := NewAttractionsStorage(db)
	res := pb.AttractionList{
		Category: "culture",
	}
	req, err := atts.ListAttractions(&res)
	if err != nil {
		t.Errorf("Failed to list attractions: %v", err)
		return
	}
	fmt.Println(req)
}

func TestSearchAttractions(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()
	atts := NewAttractionsStorage(db)
	res := pb.AttractionSearch{
		Limit:  "1",
		Offset: "0",
	}
	req, err := atts.SearchAttractions(&res)
	if err != nil {
		t.Errorf("Failed to search attractions: %v", err)
		return
	}
	fmt.Println(req)
}

func TestAddImageUrl(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()
	atts := NewAttractionsStorage(db)

	res := pb.AttractionImage{
		Id:       "af35d7ee-6bf1-4f2b-981e-f49da6b2f266",
		ImageUrl: "2222",
	}
	req, err := atts.AddImageUrl(&res)
	if err != nil {
		t.Errorf("Failed to add image url: %v", err)
		return
	}
	fmt.Println(req)
}

func TestDeleteAttraction(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()
	atts := NewAttractionsStorage(db)
	res := pb.AttractionId{
		Id: "fd2ceee5-e9c9-42d1-9f35-2488d3fb3f30",
	}
	req, err := atts.DeleteAttraction(&res)
	if err != nil {
		t.Errorf("Failed to delete attraction: %v", err)
		return
	}
	fmt.Println(req)
}

func TestRemoveHistoricalImage(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()
	atts := NewAttractionsStorage(db)
	res := pb.HistoricalImage{
		Id: "af35d7ee-6bf1-4f2b-981e-f49da6b2f266",
	}
	req, err := atts.RemoveHistoricalImage(&res)
	if err != nil {
		t.Errorf("Failed to remove historical image: %v", err)
		return
	}
	fmt.Println(req)
}
