package main

import (
	"fmt"
	"wb0-app/cache"
	"wb0-app/client"
	"wb0-app/config"
	"wb0-app/logger"
	"wb0-app/models"
	"wb0-app/storage"
)

func main() {
	fmt.Println("App run!")
	config.Load()

	logger.Setup()

	storage := storage.New()

	cache := cache.New()

	subscriber := client.New()
	subscriber.Subscribe(models.Order{}, storage.Save, cache.Save)

	select {}

}
