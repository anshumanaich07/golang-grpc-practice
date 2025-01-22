package main

import (
	grpc "learn-grpc/internal/api/grpc/server"
	"learn-grpc/internal/config"
	"log"
)

func main() {
	// read grpc configs
	grpcConf, err := config.ReadgRPCConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = grpc.RegisterService(grpcConf)
	if err != nil {
		log.Fatal(err)
	}
}
