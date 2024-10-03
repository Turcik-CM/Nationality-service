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
	loggers := logger.InitLogger()
	cfg := config.Load()

	db, err := postgres.ConnectPostgres(cfg)
	if err != nil {
		loggers.Error("Failed to connect to database")
		log.Fatal(err)
	}

	attSt := postgres.NewAttractionsStorage(db)
	hisSt := postgres.NewHistoryStorage(db)
	natSt := postgres.NewNationalFoodsStorage(db)
	cttST := postgres.NewCountriesStorage(db)
	tourismSR := service.NewTourismService(hisSt, natSt, attSt, cttST, loggers)

	listen, err := net.Listen("tcp", cfg.NATIONAL_PORT)
	fmt.Println("Listening on " + cfg.NATIONAL_PORT)
	if err != nil {
		loggers.Error("Failed to listen to tourism")
		log.Fatal(err)
	}

	server := grpc.NewServer()
	nationality.RegisterNationalityServiceServer(server, tourismSR)

	if err := server.Serve(listen); err != nil {
		loggers.Error("Error starting server on port "+cfg.NATIONAL_PORT, "error", err)
		log.Fatal(err)
	}
}
