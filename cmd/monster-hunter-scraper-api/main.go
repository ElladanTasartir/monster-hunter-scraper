package main

import (
	"context"
	config "github.com/ElladanTasartir/monster-hunter-scraper/internal/config"
	"github.com/ElladanTasartir/monster-hunter-scraper/internal/http"
	"github.com/ElladanTasartir/monster-hunter-scraper/internal/storage"
	"log"
	"time"
)

func main() {
	appConfig, err := config.NewConfig("./config")
	if err != nil {
		log.Fatalf("an error has ocurred while parsing config. err = %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	databaseStorage, err := storage.NewMongoClient(ctx, appConfig.MongoUri, appConfig.MongoDatabase)
	if err != nil {
		log.Fatalf("an error has ocurred while connecting to DB. err = %v", err)
	}

	defer func() {
		err = databaseStorage.Disconnect(ctx)
		if err != nil {
			log.Fatalf("an error has ocurred while disconnecting from DB. err = %v", err)
		}
	}()

	server, err := http.NewServer(appConfig, databaseStorage)
	if err != nil {
		log.Fatalf("an error has ocurred while creating server. err = %v", err)
	}

	err = server.Run()
	if err != nil {
		log.Fatalf("an error has ocurred while running server. err = %v", err)
	}
}
