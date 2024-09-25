package main

import (
	"fmt"
	"google.golang.org/grpc"

	"log"
	"nationality-service/genproto/nationality"
	"nationality-service/pkg/config"
	"nationality-service/pkg/logger"
	"nationality-service/service"
	"nationality-service/storage/postgres"
	"net"
)

func main() {
	logger := logger.InitLogger()
	cfg := config.Load()

	db, err := postgres.ConnectPostgres(cfg)
	if err != nil {
		logger.Error("Failed to connect to database")
		log.Fatal(err)
	}

	attSt := postgres.NewAttractionsStorage(db)
	hisSt := postgres.NewHistoryStorage(db)
	natSt := postgres.NewNationalFoodsStorage(db)
	tourismSR := service.NewTourismService(hisSt, natSt, attSt, logger)

	listen, err := net.Listen("tcp", cfg.NATIONALITY_SERVICE)
	fmt.Println("Listening on " + cfg.NATIONALITY_SERVICE)
	if err != nil {
		logger.Error("Failed to listen to tourism")
		log.Fatal(err)
	}

	server := grpc.NewServer()
	nationality.RegisterNationalityServiceServer(server, tourismSR)

	if err := server.Serve(listen); err != nil {
		logger.Error("Error starting server on port "+cfg.NATIONALITY_SERVICE, "error", err)
		log.Fatal(err)
	}
}
