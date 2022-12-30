package main

import (
	"context"
	"log"

	"github.com/takokun778/2022/internal/adapter/gateway"
	"github.com/takokun778/2022/internal/driver/config"
	"github.com/takokun778/2022/internal/driver/database"
)

func main() {
	config.Init()

	db := database.NewClient()

	rdb, err := db.Of(config.Get().DSN)
	if err != nil {
		log.Fatal(err)
	}

	tagGateway := gateway.NewTag(rdb)

	ctx := context.Background()

	if err := tagGateway.CreateTable(ctx); err != nil {
		log.Fatal(err)
	}

	if err := tagGateway.CreateIndex(ctx); err != nil {
		log.Fatal(err)
	}
}
