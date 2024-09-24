package storage

import (
	"context"
	pb "nationality-service/genproto/nationality"
)

// AttractionsStorage defines the operations related to attractions.
type AttractionsStorage interface {
	CreateAttraction(ctx context.Context, in *pb.Attraction) (*pb.AttractionResponse, error)
	GetAttractionByID(ctx context.Context, in *pb.AttractionId) (*pb.AttractionResponse, error)
	UpdateAttraction(ctx context.Context, in *pb.UpdateAttraction) (*pb.AttractionResponse, error)
	DeleteAttraction(ctx context.Context, in *pb.AttractionId) (*pb.Message, error)
	ListAttractions(ctx context.Context, in *pb.AttractionList) (*pb.AttractionListResponse, error)
	SearchAttractions(ctx context.Context, in *pb.AttractionSearch) (*pb.AttractionListResponse, error)
	AddImageUrl(ctx context.Context, in *pb.AttractionImage) (*pb.AttractionResponse, error)
}

// HistoryStorage defines the operations related to historical entities.
type HistoryStorage interface {
	AddHistorical(ctx context.Context, in *pb.Historical) (*pb.HistoricalResponse, error)
	UpdateHistoricals(ctx context.Context, in *pb.UpdateHistorical) (*pb.HistoricalResponse, error)
	GetHistoricalByID(ctx context.Context, in *pb.HistoricalId) (*pb.HistoricalResponse, error)
	DeleteHistorical(ctx context.Context, in *pb.HistoricalId) (*pb.Message, error)
	ListHistorical(ctx context.Context, in *pb.HistoricalList) (*pb.HistoricalListResponse, error)
	SearchHistorical(ctx context.Context, in *pb.HistoricalSearch) (*pb.HistoricalListResponse, error)
	AddHistoricalImage(ctx context.Context, in *pb.HistoricalImage) (*pb.Message, error)
}

// NationalFoodsStorage defines the operations related to national foods.
type NationalFoodsStorage interface {
	CreateNationalFood(ctx context.Context, in *pb.NationalFood) (*pb.NationalFoodResponse, error)
	UpdateNationalFood(ctx context.Context, in *pb.UpdateNationalFood) (*pb.NationalFoodResponse, error)
	GetNationalFoodByID(ctx context.Context, in *pb.NationalFoodId) (*pb.NationalFoodResponse, error)
	DeleteNationalFood(ctx context.Context, in *pb.NationalFoodId) (*pb.Message, error)
	ListNationalFoods(ctx context.Context, in *pb.NationalFoodList) (*pb.NationalFoodListResponse, error)
	//SearchNationalFoods(ctx context.Context, in *pb.NationalFoodSearch) (*pb.NationalFoodListResponse, error)
}
