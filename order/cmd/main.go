package main

// https://github.com/huseyinbabal/microservices/tree/main/order

import (
	"log"
	"order/config"
	"order/internal/adapters/db"
	"order/internal/adapters/grpc"
	"order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDatasourceURL())
	if err != nil {
		log.Fatalf("failed to connect to database, error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.New(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
