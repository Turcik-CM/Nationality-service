package storage

import (
	pb "nationality-service/genproto/nationality"
)

// AttractionsStorage defines the operations related to attractions.
type AttractionsStorage interface {
	CreateAttraction(in *pb.Attraction) (*pb.AttractionResponse, error)
	GetAttractionByID(in *pb.AttractionId) (*pb.AttractionResponse, error)
	UpdateAttraction(in *pb.UpdateAttraction) (*pb.AttractionResponse, error)
	DeleteAttraction(in *pb.AttractionId) (*pb.Message, error)
	ListAttractions(in *pb.AttractionList) (*pb.AttractionListResponse, error)
	SearchAttractions(in *pb.AttractionSearch) (*pb.AttractionListResponse, error)
	AddImageUrl(in *pb.AttractionImage) (*pb.Message, error)
	RemoveHistoricalImage(in *pb.HistoricalImage) (*pb.Message, error)
	CreateAttractionType(in *pb.CreateAttractionTypeRequest) (*pb.CreateAttractionTypeResponse, error)
	GetAttractionTypeByID(in *pb.GetAttractionTypeRequest) (*pb.GetAttractionTypeResponse, error)
	UpdateAttractionType(in *pb.UpdateAttractionTypeRequest) (*pb.UpdateAttractionTypeResponse, error)
	DeleteAttractionType(in *pb.DeleteAttractionTypeRequest) (*pb.Message, error)
	ListAttractionTypes(in *pb.ListAttractionTypesRequest) (*pb.ListAttractionTypesResponse, error)
}

// HistoryStorage defines the operations related to historical entities.
type HistoryStorage interface {
	AddHistorical(in *pb.Historical) (*pb.HistoricalResponse, error)
	UpdateHistoricals(in *pb.UpdateHistorical) (*pb.HistoricalResponse, error)
	GetHistoricalByID(in *pb.HistoricalId) (*pb.HistoricalResponse, error)
	DeleteHistorical(in *pb.HistoricalId) (*pb.Message, error)
	ListHistorical(in *pb.HistoricalList) (*pb.HistoricalListResponse, error)
	SearchHistorical(in *pb.HistoricalSearch) (*pb.HistoricalListResponse, error)
	AddHistoricalImage(in *pb.HistoricalImage) (*pb.Message, error)
}

// NationalFoodsStorage defines the operations related to national foods.
type NationalFoodsStorage interface {
	CreateNationalFood(in *pb.NationalFood) (*pb.NationalFoodResponse, error)
	UpdateNationalFood(in *pb.UpdateNationalFood) (*pb.NationalFoodResponse, error)
	GetNationalFoodByID(in *pb.NationalFoodId) (*pb.NationalFoodResponse, error)
	DeleteNationalFood(in *pb.NationalFoodId) (*pb.Message, error)
	ListNationalFoods(in *pb.NationalFoodList) (*pb.NationalFoodListResponse, error)
	AddImageUrll(in *pb.NationalFoodImage) (*pb.Message, error)
	//SearchNationalFoods(in *pb.NationalFoodSearch) (*pb.NationalFoodListResponse, error)
}

type CountriesStorage interface {
	CreateCountry(in *pb.CreateCountryRequest) (*pb.CreateCountryResponse, error)
	GetCountry(in *pb.GetCountryRequest) (*pb.GetCountryResponse, error)
	UpdateCountry(in *pb.UpdateCountryRequest) (*pb.UpdateCountryResponse, error)
	DeleteCountry(in *pb.DeleteCountryRequest) (*pb.Message, error)
	ListCountries(in *pb.ListCountriesRequest) (*pb.ListCountriesResponse, error)

	CreateCity(in *pb.CreateCityRequest) (*pb.CreateCityResponse, error)
	GetCity(in *pb.GetCityRequest) (*pb.CreateCityResponse, error)
	UpdateCity(in *pb.CreateCityResponse) (*pb.CreateCityResponse, error)
	DeleteCity(in *pb.GetCityRequest) (*pb.Message, error)
	ListCity(in *pb.ListCityRequest) (*pb.ListCityResponse, error)
	GetBYCount(in *pb.CountryId) (*pb.GetCountryId, error)
}
