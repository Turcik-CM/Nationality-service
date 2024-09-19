package main

import (
	"google.golang.org/grpc

	"log"
	"net"
	"tourism-service/pkg/config"
	"tourism-service/pkg/logger"
	"tourism-service/service"
	"tourism-service/storage/postgres"
)

func main() {
	logger := logger.InitLogger()
	cfg := config.Load()

	db, err := postgres.ConnectPostgres(cfg)
	if err != nil {
		logger.Error("Failed to connect to database")
		log.Fatal(err)
	}

	tourismSt := postgres.NewHistoryRepo(db)
	tourismSR := service.NewTourismService(tourismSt)

	listen, err := net.Listen("tcp", cfg.TOURISM_SERVICE)
	if err != nil {
		logger.Error("Failed to listen to tourism")
		log.Fatal(err)
	}

	service := grpc.NewServer()

}
