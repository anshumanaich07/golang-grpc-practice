package main

import (
	grpc "learn-grpc/internal/api/grpc/server"
	"learn-grpc/internal/config"
	"learn-grpc/internal/database"
	"log"
)

func main() {
	// read grpc configs
	grpcConf, err := config.ReadgRPCConfig()
	if err != nil {
		log.Fatal(err)
	}
	dbConfig, err := config.ReadDBConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.GetDB(*dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	// registering grpc services
	err = grpc.RegisterUserService(grpcConf, db)
	if err != nil {
		log.Fatal(err)
	}
}
